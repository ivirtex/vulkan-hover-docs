# VkWriteDescriptorSetAccelerationStructureNV(3) Manual Page

## Name

VkWriteDescriptorSetAccelerationStructureNV - Structure specifying
acceleration structure descriptor information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkWriteDescriptorSetAccelerationStructureNV` structure is defined
as:

``` c
// Provided by VK_NV_ray_tracing
typedef struct VkWriteDescriptorSetAccelerationStructureNV {
    VkStructureType                     sType;
    const void*                         pNext;
    uint32_t                            accelerationStructureCount;
    const VkAccelerationStructureNV*    pAccelerationStructures;
} VkWriteDescriptorSetAccelerationStructureNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `accelerationStructureCount` is the number of elements in
  `pAccelerationStructures`.

- `pAccelerationStructures` is a pointer to an array of
  [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html) structures
  specifying the acceleration structures to update.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkWriteDescriptorSetAccelerationStructureNV-accelerationStructureCount-03747"
  id="VUID-VkWriteDescriptorSetAccelerationStructureNV-accelerationStructureCount-03747"></a>
  VUID-VkWriteDescriptorSetAccelerationStructureNV-accelerationStructureCount-03747  
  `accelerationStructureCount` **must** be equal to `descriptorCount` in
  the extended structure

- <a
  href="#VUID-VkWriteDescriptorSetAccelerationStructureNV-pAccelerationStructures-03748"
  id="VUID-VkWriteDescriptorSetAccelerationStructureNV-pAccelerationStructures-03748"></a>
  VUID-VkWriteDescriptorSetAccelerationStructureNV-pAccelerationStructures-03748  
  Each acceleration structure in `pAccelerationStructures` **must** have
  been created with `VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR`

- <a
  href="#VUID-VkWriteDescriptorSetAccelerationStructureNV-pAccelerationStructures-03749"
  id="VUID-VkWriteDescriptorSetAccelerationStructureNV-pAccelerationStructures-03749"></a>
  VUID-VkWriteDescriptorSetAccelerationStructureNV-pAccelerationStructures-03749  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nullDescriptor"
  target="_blank" rel="noopener"><code>nullDescriptor</code></a> feature
  is not enabled, each member of `pAccelerationStructures` **must** not
  be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

Valid Usage (Implicit)

- <a href="#VUID-VkWriteDescriptorSetAccelerationStructureNV-sType-sType"
  id="VUID-VkWriteDescriptorSetAccelerationStructureNV-sType-sType"></a>
  VUID-VkWriteDescriptorSetAccelerationStructureNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_ACCELERATION_STRUCTURE_NV`

- <a
  href="#VUID-VkWriteDescriptorSetAccelerationStructureNV-pAccelerationStructures-parameter"
  id="VUID-VkWriteDescriptorSetAccelerationStructureNV-pAccelerationStructures-parameter"></a>
  VUID-VkWriteDescriptorSetAccelerationStructureNV-pAccelerationStructures-parameter  
  `pAccelerationStructures` **must** be a valid pointer to an array of
  `accelerationStructureCount` valid or
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)
  [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html) handles

- <a
  href="#VUID-VkWriteDescriptorSetAccelerationStructureNV-accelerationStructureCount-arraylength"
  id="VUID-VkWriteDescriptorSetAccelerationStructureNV-accelerationStructureCount-arraylength"></a>
  VUID-VkWriteDescriptorSetAccelerationStructureNV-accelerationStructureCount-arraylength  
  `accelerationStructureCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkWriteDescriptorSetAccelerationStructureNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
