# vkGetCalibratedTimestampsKHR(3) Manual Page

## Name

vkGetCalibratedTimestampsKHR - Query calibrated timestamps



## <a href="#_c_specification" class="anchor"></a>C Specification

In order to be able to correlate the time a particular operation took
place at on timelines of different time domains (e.g. a device operation
vs. a host operation), Vulkan allows querying calibrated timestamps from
multiple time domains.

To query calibrated timestamps from a set of time domains, call:

``` c
// Provided by VK_KHR_calibrated_timestamps
VkResult vkGetCalibratedTimestampsKHR(
    VkDevice                                    device,
    uint32_t                                    timestampCount,
    const VkCalibratedTimestampInfoKHR*         pTimestampInfos,
    uint64_t*                                   pTimestamps,
    uint64_t*                                   pMaxDeviation);
```

or the equivalent command

``` c
// Provided by VK_EXT_calibrated_timestamps
VkResult vkGetCalibratedTimestampsEXT(
    VkDevice                                    device,
    uint32_t                                    timestampCount,
    const VkCalibratedTimestampInfoKHR*         pTimestampInfos,
    uint64_t*                                   pTimestamps,
    uint64_t*                                   pMaxDeviation);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device used to perform the query.

- `timestampCount` is the number of timestamps to query.

- `pTimestampInfos` is a pointer to an array of `timestampCount`
  [VkCalibratedTimestampInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCalibratedTimestampInfoKHR.html)
  structures, describing the time domains the calibrated timestamps
  should be captured from.

- `pTimestamps` is a pointer to an array of `timestampCount` 64-bit
  unsigned integer values in which the requested calibrated timestamp
  values are returned.

- `pMaxDeviation` is a pointer to a 64-bit unsigned integer value in
  which the strictly positive maximum deviation, in nanoseconds, of the
  calibrated timestamp values is returned.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The maximum deviation <strong>may</strong> vary between calls to
<code>vkGetCalibratedTimestampsKHR</code> even for the same set of time
domains due to implementation and platform specific reasons. It is the
application’s responsibility to assess whether the returned maximum
deviation makes the timestamp values suitable for any particular purpose
and <strong>can</strong> choose to re-issue the timestamp calibration
call pursuing a lower deviation value.</p></td>
</tr>
</tbody>
</table>

Calibrated timestamp values **can** be extrapolated to estimate future
coinciding timestamp values, however, depending on the nature of the
time domains and other properties of the platform extrapolating values
over a sufficiently long period of time **may** no longer be accurate
enough to fit any particular purpose, so applications are expected to
re-calibrate the timestamps on a regular basis.

Valid Usage

- <a href="#VUID-vkGetCalibratedTimestampsEXT-timeDomain-09246"
  id="VUID-vkGetCalibratedTimestampsEXT-timeDomain-09246"></a>
  VUID-vkGetCalibratedTimestampsEXT-timeDomain-09246  
  The `timeDomain` value of each
  [VkCalibratedTimestampInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCalibratedTimestampInfoKHR.html) in
  `pTimestampInfos` **must** be unique

Valid Usage (Implicit)

- <a href="#VUID-vkGetCalibratedTimestampsKHR-device-parameter"
  id="VUID-vkGetCalibratedTimestampsKHR-device-parameter"></a>
  VUID-vkGetCalibratedTimestampsKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetCalibratedTimestampsKHR-pTimestampInfos-parameter"
  id="VUID-vkGetCalibratedTimestampsKHR-pTimestampInfos-parameter"></a>
  VUID-vkGetCalibratedTimestampsKHR-pTimestampInfos-parameter  
  `pTimestampInfos` **must** be a valid pointer to an array of
  `timestampCount` valid
  [VkCalibratedTimestampInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCalibratedTimestampInfoKHR.html)
  structures

- <a href="#VUID-vkGetCalibratedTimestampsKHR-pTimestamps-parameter"
  id="VUID-vkGetCalibratedTimestampsKHR-pTimestamps-parameter"></a>
  VUID-vkGetCalibratedTimestampsKHR-pTimestamps-parameter  
  `pTimestamps` **must** be a valid pointer to an array of
  `timestampCount` `uint64_t` values

- <a href="#VUID-vkGetCalibratedTimestampsKHR-pMaxDeviation-parameter"
  id="VUID-vkGetCalibratedTimestampsKHR-pMaxDeviation-parameter"></a>
  VUID-vkGetCalibratedTimestampsKHR-pMaxDeviation-parameter  
  `pMaxDeviation` **must** be a valid pointer to a `uint64_t` value

- <a href="#VUID-vkGetCalibratedTimestampsKHR-timestampCount-arraylength"
  id="VUID-vkGetCalibratedTimestampsKHR-timestampCount-arraylength"></a>
  VUID-vkGetCalibratedTimestampsKHR-timestampCount-arraylength  
  `timestampCount` **must** be greater than `0`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_calibrated_timestamps](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_calibrated_timestamps.html),
[VK_KHR_calibrated_timestamps](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_calibrated_timestamps.html),
[VkCalibratedTimestampInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCalibratedTimestampInfoKHR.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetCalibratedTimestampsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
