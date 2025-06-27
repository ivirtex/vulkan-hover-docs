# VkBindAccelerationStructureMemoryInfoNV(3) Manual Page

## Name

VkBindAccelerationStructureMemoryInfoNV - Structure specifying
acceleration structure memory binding



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBindAccelerationStructureMemoryInfoNV` structure is defined as:

``` c
// Provided by VK_NV_ray_tracing
typedef struct VkBindAccelerationStructureMemoryInfoNV {
    VkStructureType              sType;
    const void*                  pNext;
    VkAccelerationStructureNV    accelerationStructure;
    VkDeviceMemory               memory;
    VkDeviceSize                 memoryOffset;
    uint32_t                     deviceIndexCount;
    const uint32_t*              pDeviceIndices;
} VkBindAccelerationStructureMemoryInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `accelerationStructure` is the acceleration structure to be attached
  to memory.

- `memory` is a `VkDeviceMemory` object describing the device memory to
  attach.

- `memoryOffset` is the start offset of the region of memory that is to
  be bound to the acceleration structure. The number of bytes returned
  in the [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html)::`size`
  member in `memory`, starting from `memoryOffset` bytes, will be bound
  to the specified acceleration structure.

- `deviceIndexCount` is the number of elements in `pDeviceIndices`.

- `pDeviceIndices` is a pointer to an array of device indices.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkBindAccelerationStructureMemoryInfoNV-accelerationStructure-03620"
  id="VUID-VkBindAccelerationStructureMemoryInfoNV-accelerationStructure-03620"></a>
  VUID-VkBindAccelerationStructureMemoryInfoNV-accelerationStructure-03620  
  `accelerationStructure` **must** not already be backed by a memory
  object

- <a
  href="#VUID-VkBindAccelerationStructureMemoryInfoNV-memoryOffset-03621"
  id="VUID-VkBindAccelerationStructureMemoryInfoNV-memoryOffset-03621"></a>
  VUID-VkBindAccelerationStructureMemoryInfoNV-memoryOffset-03621  
  `memoryOffset` **must** be less than the size of `memory`

- <a href="#VUID-VkBindAccelerationStructureMemoryInfoNV-memory-03622"
  id="VUID-VkBindAccelerationStructureMemoryInfoNV-memory-03622"></a>
  VUID-VkBindAccelerationStructureMemoryInfoNV-memory-03622  
  `memory` **must** have been allocated using one of the memory types
  allowed in the `memoryTypeBits` member of the
  [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html) structure returned
  from a call to
  [vkGetAccelerationStructureMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureMemoryRequirementsNV.html)
  with `accelerationStructure` and `type` of
  `VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_OBJECT_NV`

- <a
  href="#VUID-VkBindAccelerationStructureMemoryInfoNV-memoryOffset-03623"
  id="VUID-VkBindAccelerationStructureMemoryInfoNV-memoryOffset-03623"></a>
  VUID-VkBindAccelerationStructureMemoryInfoNV-memoryOffset-03623  
  `memoryOffset` **must** be an integer multiple of the `alignment`
  member of the [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html)
  structure returned from a call to
  [vkGetAccelerationStructureMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureMemoryRequirementsNV.html)
  with `accelerationStructure` and `type` of
  `VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_OBJECT_NV`

- <a href="#VUID-VkBindAccelerationStructureMemoryInfoNV-size-03624"
  id="VUID-VkBindAccelerationStructureMemoryInfoNV-size-03624"></a>
  VUID-VkBindAccelerationStructureMemoryInfoNV-size-03624  
  The `size` member of the `VkMemoryRequirements` structure returned
  from a call to
  [vkGetAccelerationStructureMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureMemoryRequirementsNV.html)
  with `accelerationStructure` and `type` of
  `VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_OBJECT_NV`
  **must** be less than or equal to the size of `memory` minus
  `memoryOffset`

Valid Usage (Implicit)

- <a href="#VUID-VkBindAccelerationStructureMemoryInfoNV-sType-sType"
  id="VUID-VkBindAccelerationStructureMemoryInfoNV-sType-sType"></a>
  VUID-VkBindAccelerationStructureMemoryInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_BIND_ACCELERATION_STRUCTURE_MEMORY_INFO_NV`

- <a href="#VUID-VkBindAccelerationStructureMemoryInfoNV-pNext-pNext"
  id="VUID-VkBindAccelerationStructureMemoryInfoNV-pNext-pNext"></a>
  VUID-VkBindAccelerationStructureMemoryInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`

- <a
  href="#VUID-VkBindAccelerationStructureMemoryInfoNV-accelerationStructure-parameter"
  id="VUID-VkBindAccelerationStructureMemoryInfoNV-accelerationStructure-parameter"></a>
  VUID-VkBindAccelerationStructureMemoryInfoNV-accelerationStructure-parameter  
  `accelerationStructure` **must** be a valid
  [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html) handle

- <a href="#VUID-VkBindAccelerationStructureMemoryInfoNV-memory-parameter"
  id="VUID-VkBindAccelerationStructureMemoryInfoNV-memory-parameter"></a>
  VUID-VkBindAccelerationStructureMemoryInfoNV-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)
  handle

- <a
  href="#VUID-VkBindAccelerationStructureMemoryInfoNV-pDeviceIndices-parameter"
  id="VUID-VkBindAccelerationStructureMemoryInfoNV-pDeviceIndices-parameter"></a>
  VUID-VkBindAccelerationStructureMemoryInfoNV-pDeviceIndices-parameter  
  If `deviceIndexCount` is not `0`, `pDeviceIndices` **must** be a valid
  pointer to an array of `deviceIndexCount` `uint32_t` values

- <a href="#VUID-VkBindAccelerationStructureMemoryInfoNV-commonparent"
  id="VUID-VkBindAccelerationStructureMemoryInfoNV-commonparent"></a>
  VUID-VkBindAccelerationStructureMemoryInfoNV-commonparent  
  Both of `accelerationStructure`, and `memory` **must** have been
  created, allocated, or retrieved from the same
  [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html),
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkBindAccelerationStructureMemoryNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindAccelerationStructureMemoryNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBindAccelerationStructureMemoryInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
