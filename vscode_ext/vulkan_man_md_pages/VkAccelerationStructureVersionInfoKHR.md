# VkAccelerationStructureVersionInfoKHR(3) Manual Page

## Name

VkAccelerationStructureVersionInfoKHR - Acceleration structure version
information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAccelerationStructureVersionInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_acceleration_structure
typedef struct VkAccelerationStructureVersionInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    const uint8_t*     pVersionData;
} VkAccelerationStructureVersionInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pVersionData` is a pointer to the version header of an acceleration
  structure as defined in
  [vkCmdCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyAccelerationStructureToMemoryKHR.html)

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p><code>pVersionData</code> is a <em>pointer</em> to an array of
2×<code>VK_UUID_SIZE</code> <code>uint8_t</code> values instead of two
<code>VK_UUID_SIZE</code> arrays as the expected use case for this
member is to be pointed at the header of a previously serialized
acceleration structure (via <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyAccelerationStructureToMemoryKHR.html">vkCmdCopyAccelerationStructureToMemoryKHR</a>
or <a
href="vkCopyAccelerationStructureToMemoryKHR.html">vkCopyAccelerationStructureToMemoryKHR</a>)
that is loaded in memory. Using arrays would necessitate extra memory
copies of the UUIDs.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-VkAccelerationStructureVersionInfoKHR-sType-sType"
  id="VUID-VkAccelerationStructureVersionInfoKHR-sType-sType"></a>
  VUID-VkAccelerationStructureVersionInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_VERSION_INFO_KHR`

- <a href="#VUID-VkAccelerationStructureVersionInfoKHR-pNext-pNext"
  id="VUID-VkAccelerationStructureVersionInfoKHR-pNext-pNext"></a>
  VUID-VkAccelerationStructureVersionInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a
  href="#VUID-VkAccelerationStructureVersionInfoKHR-pVersionData-parameter"
  id="VUID-VkAccelerationStructureVersionInfoKHR-pVersionData-parameter"></a>
  VUID-VkAccelerationStructureVersionInfoKHR-pVersionData-parameter  
  `pVersionData` **must** be a valid pointer to an array of
  2×VK_UUID_SIZE `uint8_t` values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetDeviceAccelerationStructureCompatibilityKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceAccelerationStructureCompatibilityKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAccelerationStructureVersionInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
