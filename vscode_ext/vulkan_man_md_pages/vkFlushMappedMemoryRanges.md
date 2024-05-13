# vkFlushMappedMemoryRanges(3) Manual Page

## Name

vkFlushMappedMemoryRanges - Flush mapped memory ranges



## <a href="#_c_specification" class="anchor"></a>C Specification

To flush ranges of non-coherent memory from the host caches, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkFlushMappedMemoryRanges(
    VkDevice                                    device,
    uint32_t                                    memoryRangeCount,
    const VkMappedMemoryRange*                  pMemoryRanges);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the memory ranges.

- `memoryRangeCount` is the length of the `pMemoryRanges` array.

- `pMemoryRanges` is a pointer to an array of
  [VkMappedMemoryRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMappedMemoryRange.html) structures describing
  the memory ranges to flush.

## <a href="#_description" class="anchor"></a>Description

`vkFlushMappedMemoryRanges` guarantees that host writes to the memory
ranges described by `pMemoryRanges` are made available to the host
memory domain, such that they **can** be made available to the device
memory domain via <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-available-and-visible"
target="_blank" rel="noopener">memory domain operations</a> using the
`VK_ACCESS_HOST_WRITE_BIT` <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-types"
target="_blank" rel="noopener">access type</a>.

Within each range described by `pMemoryRanges`, each set of
`nonCoherentAtomSize` bytes in that range is flushed if any byte in that
set has been written by the host since it was first host mapped, or the
last time it was flushed. If `pMemoryRanges` includes sets of
`nonCoherentAtomSize` bytes where no bytes have been written by the
host, those bytes **must** not be flushed.

Unmapping non-coherent memory does not implicitly flush the host mapped
memory, and host writes that have not been flushed **may** not ever be
visible to the device. However, implementations **must** ensure that
writes that have not been flushed do not become visible to any other
memory.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The above guarantee avoids a potential memory corruption in scenarios
where host writes to a mapped memory object have not been flushed before
the memory is unmapped (or freed), and the virtual address range is
subsequently reused for a different mapping (or memory
allocation).</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-vkFlushMappedMemoryRanges-device-parameter"
  id="VUID-vkFlushMappedMemoryRanges-device-parameter"></a>
  VUID-vkFlushMappedMemoryRanges-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkFlushMappedMemoryRanges-pMemoryRanges-parameter"
  id="VUID-vkFlushMappedMemoryRanges-pMemoryRanges-parameter"></a>
  VUID-vkFlushMappedMemoryRanges-pMemoryRanges-parameter  
  `pMemoryRanges` **must** be a valid pointer to an array of
  `memoryRangeCount` valid
  [VkMappedMemoryRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMappedMemoryRange.html) structures

- <a href="#VUID-vkFlushMappedMemoryRanges-memoryRangeCount-arraylength"
  id="VUID-vkFlushMappedMemoryRanges-memoryRangeCount-arraylength"></a>
  VUID-vkFlushMappedMemoryRanges-memoryRangeCount-arraylength  
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
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkFlushMappedMemoryRanges"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
