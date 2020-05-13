#include <cstdint>
#include <complex>
#include <chrono> 
#include <thread>
#include <vector>  
#include <iostream>
#include <iomanip>

using namespace std;
using uint8 = uint8_t;

/*
 * @brief   Helpful Timer
 */
class Timer
{
public:
    Timer() : beg_(clock_::now()) {}
    void reset() { beg_ = clock_::now(); }
    double elapsed() const { 
        return chrono::duration_cast<second_>(clock_::now() - beg_).count(); }

private:
    typedef chrono::high_resolution_clock clock_;
    typedef chrono::duration<double, ratio<1> > second_;
    chrono::time_point<clock_> beg_;
};

struct alignas(4) RGBA {
    uint8 r { 0 };
    uint8 g { 0 };
    uint8 b { 0 };
    uint8 a { 255 };
};

/*
 * @brief   Actual parallel mandelbrot generator
 */
class ParallelMandelbrotGenerator {
public:
    ParallelMandelbrotGenerator(size_t width, size_t height, RGBA* destImg, int xmin = -2, int ymin = -2, int xmax = 2, int ymax = 2) :
        width(width), height(height), img(destImg), xmin(xmin), ymin(ymin), xmax(xmax), ymax(ymax) {}

    void generate(size_t numParallel) {
        processParallelInSegments(numParallel);
    }

private:
    void processParallelInSegments(int numSegments) {
        vector<thread> workers;
        workers.reserve(numSegments);

        for (int nSegment = 0 ; nSegment < numSegments; nSegment++) {
            const int yBegin = nSegment*height/numSegments;
            const int yEnd = (nSegment+1)*height/numSegments;
            const auto procesSegment_ = [this, yBegin, yEnd]() { processSegment(yBegin, yEnd); };
            workers.emplace_back(procesSegment_);
        }

        for (thread& t : workers)
            t.join();
    }

    void processSegment(int beginy, int endy) {
        for (int py = beginy; py< endy; py++) {
            double y = double(py) / height * (ymax - ymin) + ymin;
            for (int px = 0; px < width; px++) {
                double x = double(px) / width * (xmax - xmin) + xmin;
                uint8 c = mandelbrot(x, y);
                img[px + py * width] = RGBA{c, c, c};
            }
        }
    }

    // slow impl based on complex128
    // uint8 mandelbrot(double x, double y) {
    //     using complex128 = complex<double>;
    //     complex128 z {x, y};
    //     const int iterations = 200;
    //     const int contrast = 15;

    //     complex128 v;
    //     for (int n = 0; n < iterations; n++) {
    //        v = v*v + z; // this complex128 operation is expensive
    //         if (abs(v) > 2.0) // this abs on complex makes total mandelbrot generation 2x longer (230ms vs 90ms single thread)
    //             return 255 - contrast * n;
    //     }
    //     return 0;
    // }

    // 2x faster impl based on primitive float64
    uint8 mandelbrot(double x, double y) {
        const int iterations = 200;
        const int contrast = 15;

        double real = 0.0;
        double imag = 0.0;

        for (int n = 0; n < iterations; n++) {
            const double i = imag;
            const double r = real;

            real = (r * r - i * i) + x;
            imag = (r * i + i * r) + y;

            if (sqrt(imag*imag+real*real) > 2.0)
                return 255 - contrast * n;
        }
        return 0;
    }

private:
    const size_t width;
    const size_t height;
    const int xmin;
    const int ymin;
    const int xmax;
    const int ymax;
    RGBA* img;
};

/*
 * @brief   C interface to mandelbrot generator
 */
extern "C" char* makeMandel(size_t width, size_t height, size_t numParallel) {
    RGBA* rgba = new RGBA[width*height];
    ParallelMandelbrotGenerator mandel(width, height, rgba);
    mandel.generate(numParallel);
    return (char*)rgba;
}

extern "C" void freeMandel(char* rgba) {
    delete[] rgba;
}

/*
 * @brief   C interface to mandelbrot generator, Version2
 */
extern "C" void makeMandel2(size_t width, size_t height, char* data, size_t numParallel) {
    ParallelMandelbrotGenerator mandel(width, height, (RGBA*)data);
    mandel.generate(numParallel);
}

/*
 * @brief   Standalone benchmarking exectutable
 *          CGO defined in cgomandel.go
 */
#ifndef CGO
void benchMandel(size_t width, size_t height) {
    cout << "c++ times for parallel mandel "<< width<< "x" << height << ":" << endl;

    Timer t;
    RGBA* rgba = new RGBA[width*height];

    for (size_t numSegments = 1; numSegments < height; numSegments *= 2) {
        t.reset();
        makeMandel2(width, height, (char*)rgba, numSegments);
        cout << setw(4) << numSegments << " - " << int(1000*t.elapsed()) << "ms" << endl;
    }
    
    delete[] rgba;
}

int main() {
    constexpr size_t size { 1*1024 };
    benchMandel(size, size);
}
#endif