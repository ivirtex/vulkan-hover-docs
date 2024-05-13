# vkInvalidateMappedMemoryRanges(3) Manual Page

## Name

vkInvalidateMappedMemoryRanges - Invalidate ranges of mapped memory
objects



## <a href="#_c_specification" class="anchor"></a>C Specification

To invalidate ranges of non-coherent memory from the host caches, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkInvalidateMappedMemoryRanges(
    VkDevice                                    device,
    uint32_t                                    memoryRangeCount,
    const VkMappedMemoryRange*                  pMemoryRanges);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the memory ranges.

- `memoryRangeCount` is the length of the `pMemoryRanges` array.

- `pMemoryRanges` is a pointer to an array of
  [VkMappedMemoryRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMappedMemoryRange.html) structures describing
  the memory ranges to invalidate.

## <a href="#_description" class="anchor"></a>Description

`vkInvalidateMappedMemoryRanges` guarantees that device writes to the
memory ranges described by `pMemoryRanges`, which have been made
available to the host memory domain using the `VK_ACCESS_HOST_WRITE_BIT`
and `VK_ACCESS_HOST_READ_BIT` <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-types"
target="_blank" rel="noopener">access types</a>, are made visible to the
host. If a range of non-coherent memory is written by the host and then
invalidated without first being flushed, its contents are undefined.

Within each range described by `pMemoryRanges`, each set of
`nonCoherentAtomSize` bytes in that range is invalidated if any byte in
that set has been written by the device since it was first host mapped,
or the last time it was invalidated.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Mapping non-coherent memory does not implicitly invalidate that
memory.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-vkInvalidateMappedMemoryRanges-device-parameter"
  id="VUID-vkInvalidateMappedMemoryRanges-device-parameter"></a>
  VUID-vkInvalidateMappedMemoryRanges-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkInvalidateMappedMemoryRanges-pMemoryRanges-parameter"
  id="VUID-vkInvalidateMappedMemoryRanges-pMemoryRanges-parameter"></a>
  VUID-vkInvalidateMappedMemoryRanges-pMemoryRanges-parameter  
  `pMemoryRanges` **must** be a valid pointer to an array of
  `memoryRangeCount` valid
  [VkMappedMemoryRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMappedMemoryRange.html) structures

- <a
  href="#VUID-vkInvalidateMappedMemoryRanges-memoryRangeCount-arraylength"
  id="VUID-vkInvalidateMappedMemoryRanges-memoryRangeCount-arraylength"></a>
  VUID-vkInvalidateMappedMemoryRanges-memoryRangeCount-arraylength  
  `memoryRangeCount` **must** be greater than `0`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkMappedMemoryRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMappedMemoryRange.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkInvalidateMappedMemoryRanges"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
