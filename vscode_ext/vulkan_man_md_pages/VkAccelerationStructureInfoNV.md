# VkAccelerationStructureInfoNV(3) Manual Page

## Name

VkAccelerationStructureInfoNV - Structure specifying the parameters of
acceleration structure object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAccelerationStructureInfoNV` structure is defined as:

``` c
// Provided by VK_NV_ray_tracing
typedef struct VkAccelerationStructureInfoNV {
    VkStructureType                        sType;
    const void*                            pNext;
    VkAccelerationStructureTypeNV          type;
    VkBuildAccelerationStructureFlagsNV    flags;
    uint32_t                               instanceCount;
    uint32_t                               geometryCount;
    const VkGeometryNV*                    pGeometries;
} VkAccelerationStructureInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `type` is a
  [VkAccelerationStructureTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureTypeNV.html)
  value specifying the type of acceleration structure that will be
  created.

- `flags` is a bitmask of
  [VkBuildAccelerationStructureFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildAccelerationStructureFlagBitsNV.html)
  specifying additional parameters of the acceleration structure.

- `instanceCount` specifies the number of instances that will be in the
  new acceleration structure.

- `geometryCount` specifies the number of geometries that will be in the
  new acceleration structure.

- `pGeometries` is a pointer to an array of `geometryCount`
  [VkGeometryNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryNV.html) structures containing the scene data
  being passed into the acceleration structure.

## <a href="#_description" class="anchor"></a>Description

`VkAccelerationStructureInfoNV` contains information that is used both
for acceleration structure creation with
[vkCreateAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateAccelerationStructureNV.html)
and in combination with the actual geometric data to build the
acceleration structure with
[vkCmdBuildAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBuildAccelerationStructureNV.html).

Valid Usage

- <a href="#VUID-VkAccelerationStructureInfoNV-geometryCount-02422"
  id="VUID-VkAccelerationStructureInfoNV-geometryCount-02422"></a>
  VUID-VkAccelerationStructureInfoNV-geometryCount-02422  
  `geometryCount` **must** be less than or equal to
  [VkPhysicalDeviceRayTracingPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPropertiesNV.html)::`maxGeometryCount`

- <a href="#VUID-VkAccelerationStructureInfoNV-instanceCount-02423"
  id="VUID-VkAccelerationStructureInfoNV-instanceCount-02423"></a>
  VUID-VkAccelerationStructureInfoNV-instanceCount-02423  
  `instanceCount` **must** be less than or equal to
  [VkPhysicalDeviceRayTracingPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPropertiesNV.html)::`maxInstanceCount`

- <a href="#VUID-VkAccelerationStructureInfoNV-maxTriangleCount-02424"
  id="VUID-VkAccelerationStructureInfoNV-maxTriangleCount-02424"></a>
  VUID-VkAccelerationStructureInfoNV-maxTriangleCount-02424  
  The total number of triangles in all geometries **must** be less than
  or equal to
  [VkPhysicalDeviceRayTracingPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPropertiesNV.html)::`maxTriangleCount`

- <a href="#VUID-VkAccelerationStructureInfoNV-type-02425"
  id="VUID-VkAccelerationStructureInfoNV-type-02425"></a>
  VUID-VkAccelerationStructureInfoNV-type-02425  
  If `type` is `VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_NV` then
  `geometryCount` **must** be `0`

- <a href="#VUID-VkAccelerationStructureInfoNV-type-02426"
  id="VUID-VkAccelerationStructureInfoNV-type-02426"></a>
  VUID-VkAccelerationStructureInfoNV-type-02426  
  If `type` is `VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_NV` then
  `instanceCount` **must** be `0`

- <a href="#VUID-VkAccelerationStructureInfoNV-type-02786"
  id="VUID-VkAccelerationStructureInfoNV-type-02786"></a>
  VUID-VkAccelerationStructureInfoNV-type-02786  
  If `type` is `VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_NV` then the
  `geometryType` member of each geometry in `pGeometries` **must** be
  the same

- <a href="#VUID-VkAccelerationStructureInfoNV-type-04623"
  id="VUID-VkAccelerationStructureInfoNV-type-04623"></a>
  VUID-VkAccelerationStructureInfoNV-type-04623  
  `type` **must** not be `VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR`

- <a href="#VUID-VkAccelerationStructureInfoNV-flags-02592"
  id="VUID-VkAccelerationStructureInfoNV-flags-02592"></a>
  VUID-VkAccelerationStructureInfoNV-flags-02592  
  If `flags` has the
  `VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_TRACE_BIT_NV` bit set,
  then it **must** not have the
  `VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_BUILD_BIT_NV` bit set

- <a href="#VUID-VkAccelerationStructureInfoNV-scratch-02781"
  id="VUID-VkAccelerationStructureInfoNV-scratch-02781"></a>
  VUID-VkAccelerationStructureInfoNV-scratch-02781  
  `scratch` **must** have been created with
  `VK_BUFFER_USAGE_RAY_TRACING_BIT_NV` usage flag

- <a href="#VUID-VkAccelerationStructureInfoNV-instanceData-02782"
  id="VUID-VkAccelerationStructureInfoNV-instanceData-02782"></a>
  VUID-VkAccelerationStructureInfoNV-instanceData-02782  
  If `instanceData` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `instanceData` **must** have been created with
  `VK_BUFFER_USAGE_RAY_TRACING_BIT_NV` usage flag

Valid Usage (Implicit)

- <a href="#VUID-VkAccelerationStructureInfoNV-sType-sType"
  id="VUID-VkAccelerationStructureInfoNV-sType-sType"></a>
  VUID-VkAccelerationStructureInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_INFO_NV`

- <a href="#VUID-VkAccelerationStructureInfoNV-pNext-pNext"
  id="VUID-VkAccelerationStructureInfoNV-pNext-pNext"></a>
  VUID-VkAccelerationStructureInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkAccelerationStructureInfoNV-type-parameter"
  id="VUID-VkAccelerationStructureInfoNV-type-parameter"></a>
  VUID-VkAccelerationStructureInfoNV-type-parameter  
  `type` **must** be a valid
  [VkAccelerationStructureTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureTypeNV.html)
  value

- <a href="#VUID-VkAccelerationStructureInfoNV-flags-parameter"
  id="VUID-VkAccelerationStructureInfoNV-flags-parameter"></a>
  VUID-VkAccelerationStructureInfoNV-flags-parameter  
  `flags` **must** be a valid combination of
  [VkBuildAccelerationStructureFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildAccelerationStructureFlagBitsNV.html)
  values

- <a href="#VUID-VkAccelerationStructureInfoNV-pGeometries-parameter"
  id="VUID-VkAccelerationStructureInfoNV-pGeometries-parameter"></a>
  VUID-VkAccelerationStructureInfoNV-pGeometries-parameter  
  If `geometryCount` is not `0`, `pGeometries` **must** be a valid
  pointer to an array of `geometryCount` valid
  [VkGeometryNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryNV.html) structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkAccelerationStructureCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoNV.html),
[VkAccelerationStructureTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureTypeNV.html),
[VkBuildAccelerationStructureFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildAccelerationStructureFlagsNV.html),
[VkGeometryNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdBuildAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBuildAccelerationStructureNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAccelerationStructureInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
