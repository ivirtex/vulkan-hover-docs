# VkWriteDescriptorSetAccelerationStructureKHR(3) Manual Page

## Name

VkWriteDescriptorSetAccelerationStructureKHR - Structure specifying
acceleration structure descriptor information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkWriteDescriptorSetAccelerationStructureKHR` structure is defined
as:

``` c
// Provided by VK_KHR_acceleration_structure
typedef struct VkWriteDescriptorSetAccelerationStructureKHR {
    VkStructureType                      sType;
    const void*                          pNext;
    uint32_t                             accelerationStructureCount;
    const VkAccelerationStructureKHR*    pAccelerationStructures;
} VkWriteDescriptorSetAccelerationStructureKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `accelerationStructureCount` is the number of elements in
  `pAccelerationStructures`.

- `pAccelerationStructures` is a pointer to an array of
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html)
  structures specifying the acceleration structures to update.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkWriteDescriptorSetAccelerationStructureKHR-accelerationStructureCount-02236"
  id="VUID-VkWriteDescriptorSetAccelerationStructureKHR-accelerationStructureCount-02236"></a>
  VUID-VkWriteDescriptorSetAccelerationStructureKHR-accelerationStructureCount-02236  
  `accelerationStructureCount` **must** be equal to `descriptorCount` in
  the extended structure

- <a
  href="#VUID-VkWriteDescriptorSetAccelerationStructureKHR-pAccelerationStructures-03579"
  id="VUID-VkWriteDescriptorSetAccelerationStructureKHR-pAccelerationStructures-03579"></a>
  VUID-VkWriteDescriptorSetAccelerationStructureKHR-pAccelerationStructures-03579  
  Each acceleration structure in `pAccelerationStructures` **must** have
  been created with a `type` of
  `VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR` or
  `VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR`

- <a
  href="#VUID-VkWriteDescriptorSetAccelerationStructureKHR-pAccelerationStructures-03580"
  id="VUID-VkWriteDescriptorSetAccelerationStructureKHR-pAccelerationStructures-03580"></a>
  VUID-VkWriteDescriptorSetAccelerationStructureKHR-pAccelerationStructures-03580  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nullDescriptor"
  target="_blank" rel="noopener"><code>nullDescriptor</code></a> feature
  is not enabled, each element of `pAccelerationStructures` **must** not
  be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

Valid Usage (Implicit)

- <a href="#VUID-VkWriteDescriptorSetAccelerationStructureKHR-sType-sType"
  id="VUID-VkWriteDescriptorSetAccelerationStructureKHR-sType-sType"></a>
  VUID-VkWriteDescriptorSetAccelerationStructureKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_ACCELERATION_STRUCTURE_KHR`

- <a
  href="#VUID-VkWriteDescriptorSetAccelerationStructureKHR-pAccelerationStructures-parameter"
  id="VUID-VkWriteDescriptorSetAccelerationStructureKHR-pAccelerationStructures-parameter"></a>
  VUID-VkWriteDescriptorSetAccelerationStructureKHR-pAccelerationStructures-parameter  
  `pAccelerationStructures` **must** be a valid pointer to an array of
  `accelerationStructureCount` valid or
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html) handles

- <a
  href="#VUID-VkWriteDescriptorSetAccelerationStructureKHR-accelerationStructureCount-arraylength"
  id="VUID-VkWriteDescriptorSetAccelerationStructureKHR-accelerationStructureCount-arraylength"></a>
  VUID-VkWriteDescriptorSetAccelerationStructureKHR-accelerationStructureCount-arraylength  
  `accelerationStructureCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkWriteDescriptorSetAccelerationStructureKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
