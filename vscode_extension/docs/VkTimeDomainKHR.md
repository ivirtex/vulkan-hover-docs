# VkTimeDomainKHR(3) Manual Page

## Name

VkTimeDomainKHR - Supported time domains



## [](#_c_specification)C Specification

The set of supported time domains consists of:

```c++
// Provided by VK_KHR_calibrated_timestamps
typedef enum VkTimeDomainKHR {
    VK_TIME_DOMAIN_DEVICE_KHR = 0,
    VK_TIME_DOMAIN_CLOCK_MONOTONIC_KHR = 1,
    VK_TIME_DOMAIN_CLOCK_MONOTONIC_RAW_KHR = 2,
    VK_TIME_DOMAIN_QUERY_PERFORMANCE_COUNTER_KHR = 3,
  // Provided by VK_EXT_calibrated_timestamps
    VK_TIME_DOMAIN_DEVICE_EXT = VK_TIME_DOMAIN_DEVICE_KHR,
  // Provided by VK_EXT_calibrated_timestamps
    VK_TIME_DOMAIN_CLOCK_MONOTONIC_EXT = VK_TIME_DOMAIN_CLOCK_MONOTONIC_KHR,
  // Provided by VK_EXT_calibrated_timestamps
    VK_TIME_DOMAIN_CLOCK_MONOTONIC_RAW_EXT = VK_TIME_DOMAIN_CLOCK_MONOTONIC_RAW_KHR,
  // Provided by VK_EXT_calibrated_timestamps
    VK_TIME_DOMAIN_QUERY_PERFORMANCE_COUNTER_EXT = VK_TIME_DOMAIN_QUERY_PERFORMANCE_COUNTER_KHR,
} VkTimeDomainKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_calibrated_timestamps
typedef VkTimeDomainKHR VkTimeDomainEXT;
```

## [](#_description)Description

- `VK_TIME_DOMAIN_DEVICE_KHR` specifies the device time domain. Timestamp values in this time domain use the same units and are comparable with device timestamp values captured using [vkCmdWriteTimestamp](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteTimestamp.html) or [vkCmdWriteTimestamp2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteTimestamp2.html) and are defined to be incrementing according to the [`timestampPeriod`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-timestampPeriod) of the device.
- `VK_TIME_DOMAIN_CLOCK_MONOTONIC_KHR` specifies the CLOCK\_MONOTONIC time domain available on POSIX platforms. Timestamp values in this time domain are in units of nanoseconds and are comparable with platform timestamp values captured using the POSIX clock\_gettime API as computed by this example:

Note

An implementation supporting `VK_KHR_calibrated_timestamps` or `VK_EXT_calibrated_timestamps` will use the same time domain for all its [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html) so that timestamp values reported for `VK_TIME_DOMAIN_DEVICE_KHR` can be matched to any timestamp captured through [vkCmdWriteTimestamp](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteTimestamp.html) or [vkCmdWriteTimestamp2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteTimestamp2.html) .

```c
struct timespec tv;
clock_gettime(CLOCK_MONOTONIC, &tv);
return tv.tv_nsec + tv.tv_sec*1000000000ull;
```

- `VK_TIME_DOMAIN_CLOCK_MONOTONIC_RAW_KHR` specifies the CLOCK\_MONOTONIC\_RAW time domain available on POSIX platforms. Timestamp values in this time domain are in units of nanoseconds and are comparable with platform timestamp values captured using the POSIX clock\_gettime API as computed by this example:

```c
struct timespec tv;
clock_gettime(CLOCK_MONOTONIC_RAW, &tv);
return tv.tv_nsec + tv.tv_sec*1000000000ull;
```

- `VK_TIME_DOMAIN_QUERY_PERFORMANCE_COUNTER_KHR` specifies the performance counter (QPC) time domain available on Windows. Timestamp values in this time domain are in the same units as those provided by the Windows QueryPerformanceCounter API and are comparable with platform timestamp values captured using that API as computed by this example:

```c
LARGE_INTEGER counter;
QueryPerformanceCounter(&counter);
return counter.QuadPart;
```

## [](#_see_also)See Also

[VK\_EXT\_calibrated\_timestamps](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_calibrated_timestamps.html), [VK\_KHR\_calibrated\_timestamps](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_calibrated_timestamps.html), [VkCalibratedTimestampInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCalibratedTimestampInfoKHR.html), [vkGetPhysicalDeviceCalibrateableTimeDomainsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceCalibrateableTimeDomainsKHR.html), [vkGetPhysicalDeviceCalibrateableTimeDomainsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceCalibrateableTimeDomainsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTimeDomainKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0