# VkPhysicalDeviceRayTracingPropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceRayTracingPropertiesNV - Properties of the physical
device for ray tracing



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceRayTracingPropertiesNV` structure is defined as:

``` c
// Provided by VK_NV_ray_tracing
typedef struct VkPhysicalDeviceRayTracingPropertiesNV {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           shaderGroupHandleSize;
    uint32_t           maxRecursionDepth;
    uint32_t           maxShaderGroupStride;
    uint32_t           shaderGroupBaseAlignment;
    uint64_t           maxGeometryCount;
    uint64_t           maxInstanceCount;
    uint64_t           maxTriangleCount;
    uint32_t           maxDescriptorSetAccelerationStructures;
} VkPhysicalDeviceRayTracingPropertiesNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `shaderGroupHandleSize` is the size in bytes of the shader header.

- <span id="limits-maxRecursionDepth"></span> `maxRecursionDepth` is the
  maximum number of levels of recursion allowed in a trace command.

- `maxShaderGroupStride` is the maximum stride in bytes allowed between
  shader groups in the shader binding table.

- `shaderGroupBaseAlignment` is the **required** alignment in bytes for
  the base of the shader binding table.

- `maxGeometryCount` is the maximum number of geometries in the bottom
  level acceleration structure.

- `maxInstanceCount` is the maximum number of instances in the top level
  acceleration structure.

- `maxTriangleCount` is the maximum number of triangles in all
  geometries in the bottom level acceleration structure.

- `maxDescriptorSetAccelerationStructures` is the maximum number of
  acceleration structure descriptors that are allowed in a descriptor
  set.

## <a href="#_description" class="anchor"></a>Description

Due to the fact that the geometry, instance, and triangle counts are
specified at acceleration structure creation as 32-bit values,
`maxGeometryCount`, `maxInstanceCount`, and `maxTriangleCount` **must**
not exceed 2<sup>32</sup>-1.

If the `VkPhysicalDeviceRayTracingPropertiesNV` structure is included in
the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Limits specified by this structure **must** match those specified with
the same name in
[VkPhysicalDeviceAccelerationStructurePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceAccelerationStructurePropertiesKHR.html)
and
[VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html).

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceRayTracingPropertiesNV-sType-sType"
  id="VUID-VkPhysicalDeviceRayTracingPropertiesNV-sType-sType"></a>
  VUID-VkPhysicalDeviceRayTracingPropertiesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_PROPERTIES_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceRayTracingPropertiesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
