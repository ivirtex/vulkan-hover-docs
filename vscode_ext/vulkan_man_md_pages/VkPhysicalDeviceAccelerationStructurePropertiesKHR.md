# VkPhysicalDeviceAccelerationStructurePropertiesKHR(3) Manual Page

## Name

VkPhysicalDeviceAccelerationStructurePropertiesKHR - Properties of the
physical device for acceleration structure



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceAccelerationStructurePropertiesKHR` structure is
defined as:

``` c
// Provided by VK_KHR_acceleration_structure
typedef struct VkPhysicalDeviceAccelerationStructurePropertiesKHR {
    VkStructureType    sType;
    void*              pNext;
    uint64_t           maxGeometryCount;
    uint64_t           maxInstanceCount;
    uint64_t           maxPrimitiveCount;
    uint32_t           maxPerStageDescriptorAccelerationStructures;
    uint32_t           maxPerStageDescriptorUpdateAfterBindAccelerationStructures;
    uint32_t           maxDescriptorSetAccelerationStructures;
    uint32_t           maxDescriptorSetUpdateAfterBindAccelerationStructures;
    uint32_t           minAccelerationStructureScratchOffsetAlignment;
} VkPhysicalDeviceAccelerationStructurePropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-maxGeometryCount"></span> `maxGeometryCount` is the
  maximum number of geometries in the bottom level acceleration
  structure.

- <span id="limits-maxInstanceCount"></span> `maxInstanceCount` is the
  maximum number of instances in the top level acceleration structure.

- <span id="limits-maxPrimitiveCount"></span> `maxPrimitiveCount` is the
  maximum number of triangles or AABBs in all geometries in the bottom
  level acceleration structure.

- <span id="limits-maxPerStageDescriptorAccelerationStructures"></span>
  `maxPerStageDescriptorAccelerationStructures` is the maximum number of
  acceleration structure bindings that **can** be accessible to a single
  shader stage in a pipeline layout. Descriptor bindings with a
  descriptor type of `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR`
  count against this limit. Only descriptor bindings in descriptor set
  layouts created without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  count against this limit.

- <span id="limits-maxPerStageDescriptorUpdateAfterBindAccelerationStructures"></span>
  `maxPerStageDescriptorUpdateAfterBindAccelerationStructures` is
  similar to `maxPerStageDescriptorAccelerationStructures` but counts
  descriptor bindings from descriptor sets created with or without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.

- <span id="limits-maxDescriptorSetAccelerationStructures"></span>
  `maxDescriptorSetAccelerationStructures` is the maximum number of
  acceleration structure descriptors that **can** be included in
  descriptor bindings in a pipeline layout across all pipeline shader
  stages and descriptor set numbers. Descriptor bindings with a
  descriptor type of `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR`
  count against this limit. Only descriptor bindings in descriptor set
  layouts created without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  count against this limit.

- <span id="limits-maxDescriptorSetUpdateAfterBindAccelerationStructures"></span>
  `maxDescriptorSetUpdateAfterBindAccelerationStructures` is similar to
  `maxDescriptorSetAccelerationStructures` but counts descriptor
  bindings from descriptor sets created with or without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.

- <span id="limits-minAccelerationStructureScratchOffsetAlignment"></span>
  `minAccelerationStructureScratchOffsetAlignment` is the minimum
  **required** alignment, in bytes, for scratch data passed in to an
  acceleration structure build command. The value **must** be a power of
  two.

## <a href="#_description" class="anchor"></a>Description

Due to the fact that the geometry, instance, and primitive counts are
specified at acceleration structure creation as 32-bit values, <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxGeometryCount"
target="_blank" rel="noopener"><code>maxGeometryCount</code></a>, <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxInstanceCount"
target="_blank" rel="noopener"><code>maxInstanceCount</code></a>, and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxPrimitiveCount"
target="_blank" rel="noopener"><code>maxPrimitiveCount</code></a>
**must** not exceed 2<sup>32</sup>-1.

If the `VkPhysicalDeviceAccelerationStructurePropertiesKHR` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Limits specified by this structure **must** match those specified with
the same name in
[VkPhysicalDeviceRayTracingPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPropertiesNV.html).

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceAccelerationStructurePropertiesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceAccelerationStructurePropertiesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceAccelerationStructurePropertiesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ACCELERATION_STRUCTURE_PROPERTIES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceAccelerationStructurePropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
