# VkCopyAccelerationStructureToMemoryInfoKHR(3) Manual Page

## Name

VkCopyAccelerationStructureToMemoryInfoKHR - Parameters for serializing
an acceleration structure



## <a href="#_c_specification" class="anchor"></a>C Specification

``` c
// Provided by VK_KHR_acceleration_structure
typedef struct VkCopyAccelerationStructureToMemoryInfoKHR {
    VkStructureType                       sType;
    const void*                           pNext;
    VkAccelerationStructureKHR            src;
    VkDeviceOrHostAddressKHR              dst;
    VkCopyAccelerationStructureModeKHR    mode;
} VkCopyAccelerationStructureToMemoryInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `src` is the source acceleration structure for the copy

- `dst` is the device or host address to memory which is the target for
  the copy

- `mode` is a
  [VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureModeKHR.html)
  value specifying additional operations to perform during the copy.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkCopyAccelerationStructureToMemoryInfoKHR-src-04959"
  id="VUID-VkCopyAccelerationStructureToMemoryInfoKHR-src-04959"></a>
  VUID-VkCopyAccelerationStructureToMemoryInfoKHR-src-04959  
  The source acceleration structure `src` **must** have been constructed
  prior to the execution of this command

- <a href="#VUID-VkCopyAccelerationStructureToMemoryInfoKHR-dst-03561"
  id="VUID-VkCopyAccelerationStructureToMemoryInfoKHR-dst-03561"></a>
  VUID-VkCopyAccelerationStructureToMemoryInfoKHR-dst-03561  
  The memory pointed to by `dst` **must** be at least as large as the
  serialization size of `src`, as reported by
  [vkWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkWriteAccelerationStructuresPropertiesKHR.html)
  or
  [vkCmdWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteAccelerationStructuresPropertiesKHR.html)
  with a query type of
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_SIZE_KHR`

- <a href="#VUID-VkCopyAccelerationStructureToMemoryInfoKHR-mode-03412"
  id="VUID-VkCopyAccelerationStructureToMemoryInfoKHR-mode-03412"></a>
  VUID-VkCopyAccelerationStructureToMemoryInfoKHR-mode-03412  
  `mode` **must** be `VK_COPY_ACCELERATION_STRUCTURE_MODE_SERIALIZE_KHR`

Valid Usage (Implicit)

- <a href="#VUID-VkCopyAccelerationStructureToMemoryInfoKHR-sType-sType"
  id="VUID-VkCopyAccelerationStructureToMemoryInfoKHR-sType-sType"></a>
  VUID-VkCopyAccelerationStructureToMemoryInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_COPY_ACCELERATION_STRUCTURE_TO_MEMORY_INFO_KHR`

- <a href="#VUID-VkCopyAccelerationStructureToMemoryInfoKHR-pNext-pNext"
  id="VUID-VkCopyAccelerationStructureToMemoryInfoKHR-pNext-pNext"></a>
  VUID-VkCopyAccelerationStructureToMemoryInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkCopyAccelerationStructureToMemoryInfoKHR-src-parameter"
  id="VUID-VkCopyAccelerationStructureToMemoryInfoKHR-src-parameter"></a>
  VUID-VkCopyAccelerationStructureToMemoryInfoKHR-src-parameter  
  `src` **must** be a valid
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html) handle

- <a
  href="#VUID-VkCopyAccelerationStructureToMemoryInfoKHR-mode-parameter"
  id="VUID-VkCopyAccelerationStructureToMemoryInfoKHR-mode-parameter"></a>
  VUID-VkCopyAccelerationStructureToMemoryInfoKHR-mode-parameter  
  `mode` **must** be a valid
  [VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureModeKHR.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html),
[VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureModeKHR.html),
[VkDeviceOrHostAddressKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceOrHostAddressKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyAccelerationStructureToMemoryKHR.html),
[vkCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyAccelerationStructureToMemoryKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCopyAccelerationStructureToMemoryInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
