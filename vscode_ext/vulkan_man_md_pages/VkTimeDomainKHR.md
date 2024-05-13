# VkTimeDomainKHR(3) Manual Page

## Name

VkTimeDomainKHR - Supported time domains



## <a href="#_c_specification" class="anchor"></a>C Specification

The set of supported time domains consists of:

``` c
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

``` c
// Provided by VK_EXT_calibrated_timestamps
typedef VkTimeDomainKHR VkTimeDomainEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_TIME_DOMAIN_DEVICE_KHR` specifies the device time domain.
  Timestamp values in this time domain use the same units and are
  comparable with device timestamp values captured using
  [vkCmdWriteTimestamp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteTimestamp.html) or
  [vkCmdWriteTimestamp2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteTimestamp2.html) and are defined to
  be incrementing according to the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-timestampPeriod"
  target="_blank" rel="noopener"><code>timestampPeriod</code></a> of the
  device.

- `VK_TIME_DOMAIN_CLOCK_MONOTONIC_KHR` specifies the CLOCK_MONOTONIC
  time domain available on POSIX platforms. Timestamp values in this
  time domain are in units of nanoseconds and are comparable with
  platform timestamp values captured using the POSIX clock_gettime API
  as computed by this example:

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>An implementation supporting <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_calibrated_timestamps.html"><code>VK_KHR_calibrated_timestamps</code></a>
or <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_calibrated_timestamps.html"><code>VK_EXT_calibrated_timestamps</code></a>
will use the same time domain for all its <a
href="VkQueue.html">VkQueue</a> so that timestamp values reported for
<code>VK_TIME_DOMAIN_DEVICE_KHR</code> can be matched to any timestamp
captured through <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteTimestamp.html">vkCmdWriteTimestamp</a> or <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteTimestamp2.html">vkCmdWriteTimestamp2</a> .</p></td>
</tr>
</tbody>
</table>

``` c
struct timespec tv;
clock_gettime(CLOCK_MONOTONIC, &tv);
return tv.tv_nsec + tv.tv_sec*1000000000ull;
```

- `VK_TIME_DOMAIN_CLOCK_MONOTONIC_RAW_KHR` specifies the
  CLOCK_MONOTONIC_RAW time domain available on POSIX platforms.
  Timestamp values in this time domain are in units of nanoseconds and
  are comparable with platform timestamp values captured using the POSIX
  clock_gettime API as computed by this example:

``` c
struct timespec tv;
clock_gettime(CLOCK_MONOTONIC_RAW, &tv);
return tv.tv_nsec + tv.tv_sec*1000000000ull;
```

- `VK_TIME_DOMAIN_QUERY_PERFORMANCE_COUNTER_KHR` specifies the
  performance counter (QPC) time domain available on Windows. Timestamp
  values in this time domain are in the same units as those provided by
  the Windows QueryPerformanceCounter API and are comparable with
  platform timestamp values captured using that API as computed by this
  example:

``` c
LARGE_INTEGER counter;
QueryPerformanceCounter(&counter);
return counter.QuadPart;
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_calibrated_timestamps](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_calibrated_timestamps.html),
[VK_KHR_calibrated_timestamps](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_calibrated_timestamps.html),
[VkCalibratedTimestampInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCalibratedTimestampInfoKHR.html),
[vkGetPhysicalDeviceCalibrateableTimeDomainsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceCalibrateableTimeDomainsEXT.html),
[vkGetPhysicalDeviceCalibrateableTimeDomainsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceCalibrateableTimeDomainsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkTimeDomainKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
