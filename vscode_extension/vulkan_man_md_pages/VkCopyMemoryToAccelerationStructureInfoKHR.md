# VkCopyMemoryToAccelerationStructureInfoKHR(3) Manual Page

## Name

VkCopyMemoryToAccelerationStructureInfoKHR - Parameters for
deserializing an acceleration structure



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkCopyMemoryToAccelerationStructureInfoKHR` structure is defined
as:

``` c
// Provided by VK_KHR_acceleration_structure
typedef struct VkCopyMemoryToAccelerationStructureInfoKHR {
    VkStructureType                       sType;
    const void*                           pNext;
    VkDeviceOrHostAddressConstKHR         src;
    VkAccelerationStructureKHR            dst;
    VkCopyAccelerationStructureModeKHR    mode;
} VkCopyMemoryToAccelerationStructureInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `src` is the device or host address to memory containing the source
  data for the copy.

- `dst` is the target acceleration structure for the copy.

- `mode` is a
  [VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureModeKHR.html)
  value specifying additional operations to perform during the copy.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkCopyMemoryToAccelerationStructureInfoKHR-src-04960"
  id="VUID-VkCopyMemoryToAccelerationStructureInfoKHR-src-04960"></a>
  VUID-VkCopyMemoryToAccelerationStructureInfoKHR-src-04960  
  The source memory pointed to by `src` **must** contain data previously
  serialized using
  [vkCmdCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyAccelerationStructureToMemoryKHR.html),
  potentially modified to relocate acceleration structure references as
  described in that command

- <a href="#VUID-VkCopyMemoryToAccelerationStructureInfoKHR-mode-03413"
  id="VUID-VkCopyMemoryToAccelerationStructureInfoKHR-mode-03413"></a>
  VUID-VkCopyMemoryToAccelerationStructureInfoKHR-mode-03413  
  `mode` **must** be
  `VK_COPY_ACCELERATION_STRUCTURE_MODE_DESERIALIZE_KHR`

- <a href="#VUID-VkCopyMemoryToAccelerationStructureInfoKHR-pInfo-03414"
  id="VUID-VkCopyMemoryToAccelerationStructureInfoKHR-pInfo-03414"></a>
  VUID-VkCopyMemoryToAccelerationStructureInfoKHR-pInfo-03414  
  The data in `src` **must** have a format compatible with the
  destination physical device as returned by
  [vkGetDeviceAccelerationStructureCompatibilityKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceAccelerationStructureCompatibilityKHR.html)

- <a href="#VUID-VkCopyMemoryToAccelerationStructureInfoKHR-dst-03746"
  id="VUID-VkCopyMemoryToAccelerationStructureInfoKHR-dst-03746"></a>
  VUID-VkCopyMemoryToAccelerationStructureInfoKHR-dst-03746  
  `dst` **must** have been created with a `size` greater than or equal
  to that used to serialize the data in `src`

Valid Usage (Implicit)

- <a href="#VUID-VkCopyMemoryToAccelerationStructureInfoKHR-sType-sType"
  id="VUID-VkCopyMemoryToAccelerationStructureInfoKHR-sType-sType"></a>
  VUID-VkCopyMemoryToAccelerationStructureInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_COPY_MEMORY_TO_ACCELERATION_STRUCTURE_INFO_KHR`

- <a href="#VUID-VkCopyMemoryToAccelerationStructureInfoKHR-pNext-pNext"
  id="VUID-VkCopyMemoryToAccelerationStructureInfoKHR-pNext-pNext"></a>
  VUID-VkCopyMemoryToAccelerationStructureInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkCopyMemoryToAccelerationStructureInfoKHR-dst-parameter"
  id="VUID-VkCopyMemoryToAccelerationStructureInfoKHR-dst-parameter"></a>
  VUID-VkCopyMemoryToAccelerationStructureInfoKHR-dst-parameter  
  `dst` **must** be a valid
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html) handle

- <a
  href="#VUID-VkCopyMemoryToAccelerationStructureInfoKHR-mode-parameter"
  id="VUID-VkCopyMemoryToAccelerationStructureInfoKHR-mode-parameter"></a>
  VUID-VkCopyMemoryToAccelerationStructureInfoKHR-mode-parameter  
  `mode` **must** be a valid
  [VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureModeKHR.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html),
[VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureModeKHR.html),
[VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceOrHostAddressConstKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdCopyMemoryToAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMemoryToAccelerationStructureKHR.html),
[vkCopyMemoryToAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyMemoryToAccelerationStructureKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCopyMemoryToAccelerationStructureInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
