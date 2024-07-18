# VkAccelerationStructureBuildGeometryInfoKHR(3) Manual Page

## Name

VkAccelerationStructureBuildGeometryInfoKHR - Structure specifying the
geometry data used to build an acceleration structure



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAccelerationStructureBuildGeometryInfoKHR` structure is defined
as:

``` c
// Provided by VK_KHR_acceleration_structure
typedef struct VkAccelerationStructureBuildGeometryInfoKHR {
    VkStructureType                                     sType;
    const void*                                         pNext;
    VkAccelerationStructureTypeKHR                      type;
    VkBuildAccelerationStructureFlagsKHR                flags;
    VkBuildAccelerationStructureModeKHR                 mode;
    VkAccelerationStructureKHR                          srcAccelerationStructure;
    VkAccelerationStructureKHR                          dstAccelerationStructure;
    uint32_t                                            geometryCount;
    const VkAccelerationStructureGeometryKHR*           pGeometries;
    const VkAccelerationStructureGeometryKHR* const*    ppGeometries;
    VkDeviceOrHostAddressKHR                            scratchData;
} VkAccelerationStructureBuildGeometryInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `type` is a
  [VkAccelerationStructureTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureTypeKHR.html)
  value specifying the type of acceleration structure being built.

- `flags` is a bitmask of
  [VkBuildAccelerationStructureFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildAccelerationStructureFlagBitsKHR.html)
  specifying additional parameters of the acceleration structure.

- `mode` is a
  [VkBuildAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildAccelerationStructureModeKHR.html)
  value specifying the type of operation to perform.

- `srcAccelerationStructure` is a pointer to an existing acceleration
  structure that is to be used to update the `dstAccelerationStructure`
  acceleration structure when `mode` is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`.

- `dstAccelerationStructure` is a pointer to the target acceleration
  structure for the build.

- `geometryCount` specifies the number of geometries that will be built
  into `dstAccelerationStructure`.

- `pGeometries` is a pointer to an array of
  [VkAccelerationStructureGeometryKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryKHR.html)
  structures.

- `ppGeometries` is a pointer to an array of pointers to
  [VkAccelerationStructureGeometryKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryKHR.html)
  structures.

- `scratchData` is the device or host address to memory that will be
  used as scratch memory for the build.

## <a href="#_description" class="anchor"></a>Description

Only one of `pGeometries` or `ppGeometries` **can** be a valid pointer,
the other **must** be `NULL`. Each element of the non-`NULL` array
describes the data used to build each acceleration structure geometry.

The index of each element of the `pGeometries` or `ppGeometries` members
of
[VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)
is used as the *geometry index* during ray traversal. The geometry index
is available in ray shaders via the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-raygeometryindex"
target="_blank" rel="noopener"><code>RayGeometryIndexKHR</code>
built-in</a>, and is <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shader-binding-table-hit-shader-indexing"
target="_blank" rel="noopener">used to determine hit and intersection
shaders executed during traversal</a>. The geometry index is available
to ray queries via the `OpRayQueryGetIntersectionGeometryIndexKHR`
instruction.

Setting `VK_BUILD_ACCELERATION_STRUCTURE_MOTION_BIT_NV` in `flags`
indicates that this build is a motion top level acceleration structure.
A motion top level uses instances of format
[VkAccelerationStructureMotionInstanceNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInstanceNV.html)
if
[VkAccelerationStructureGeometryInstancesDataKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryInstancesDataKHR.html)::`arrayOfPointers`
is `VK_FALSE`.

If
[VkAccelerationStructureGeometryInstancesDataKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryInstancesDataKHR.html)::`arrayOfPointers`
is `VK_TRUE`, the pointer for each element of the array of instance
pointers consists of 4 bits of
`VkAccelerationStructureMotionInstanceTypeNV` in the low 4 bits of the
pointer identifying the type of structure at the pointer. The device
address accessed is the value in the array with the low 4 bits set to
zero. The structure at the pointer is one of
[VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInstanceKHR.html),
[VkAccelerationStructureMatrixMotionInstanceNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMatrixMotionInstanceNV.html)
or
[VkAccelerationStructureSRTMotionInstanceNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureSRTMotionInstanceNV.html),
depending on the type value encoded in the low 4 bits.

A top level acceleration structure with either motion instances or
vertex motion in its instances **must** set
`VK_BUILD_ACCELERATION_STRUCTURE_MOTION_BIT_NV` in `flags`.

Members `srcAccelerationStructure` and `dstAccelerationStructure`
**may** be the same or different for an update operation (when `mode` is
`VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`). If they are the
same, the update happens in-place. Otherwise, the target acceleration
structure is updated and the source is not modified.

Valid Usage

- <a href="#VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-03654"
  id="VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-03654"></a>
  VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-03654  
  `type` **must** not be `VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR`

- <a
  href="#VUID-VkAccelerationStructureBuildGeometryInfoKHR-pGeometries-03788"
  id="VUID-VkAccelerationStructureBuildGeometryInfoKHR-pGeometries-03788"></a>
  VUID-VkAccelerationStructureBuildGeometryInfoKHR-pGeometries-03788  
  If `geometryCount` is not `0`, exactly one of `pGeometries` or
  `ppGeometries` **must** be a valid pointer, the other **must** be
  `NULL`

- <a href="#VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-03789"
  id="VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-03789"></a>
  VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-03789  
  If `type` is `VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR`, the
  `geometryType` member of elements of either `pGeometries` or
  `ppGeometries` **must** be `VK_GEOMETRY_TYPE_INSTANCES_KHR`

- <a href="#VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-03790"
  id="VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-03790"></a>
  VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-03790  
  If `type` is `VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR`,
  `geometryCount` **must** be `1`

- <a href="#VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-03791"
  id="VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-03791"></a>
  VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-03791  
  If `type` is `VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_KHR` the
  `geometryType` member of elements of either `pGeometries` or
  `ppGeometries` **must** not be `VK_GEOMETRY_TYPE_INSTANCES_KHR`

- <a href="#VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-03792"
  id="VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-03792"></a>
  VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-03792  
  If `type` is `VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_KHR` then
  the `geometryType` member of each geometry in either `pGeometries` or
  `ppGeometries` **must** be the same

- <a href="#VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-03793"
  id="VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-03793"></a>
  VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-03793  
  If `type` is `VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_KHR` then
  `geometryCount` **must** be less than or equal to
  [VkPhysicalDeviceAccelerationStructurePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceAccelerationStructurePropertiesKHR.html)::`maxGeometryCount`

- <a href="#VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-03794"
  id="VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-03794"></a>
  VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-03794  
  If `type` is `VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_KHR` and the
  `geometryType` member of either `pGeometries` or `ppGeometries` is
  `VK_GEOMETRY_TYPE_AABBS_KHR`, the total number of AABBs in all
  geometries **must** be less than or equal to
  [VkPhysicalDeviceAccelerationStructurePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceAccelerationStructurePropertiesKHR.html)::`maxPrimitiveCount`

- <a href="#VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-03795"
  id="VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-03795"></a>
  VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-03795  
  If `type` is `VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_KHR` and the
  `geometryType` member of either `pGeometries` or `ppGeometries` is
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, the total number of triangles in all
  geometries **must** be less than or equal to
  [VkPhysicalDeviceAccelerationStructurePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceAccelerationStructurePropertiesKHR.html)::`maxPrimitiveCount`

- <a href="#VUID-VkAccelerationStructureBuildGeometryInfoKHR-flags-03796"
  id="VUID-VkAccelerationStructureBuildGeometryInfoKHR-flags-03796"></a>
  VUID-VkAccelerationStructureBuildGeometryInfoKHR-flags-03796  
  If `flags` has the
  `VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_TRACE_BIT_KHR` bit set,
  then it **must** not have the
  `VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_BUILD_BIT_KHR` bit set

- <a
  href="#VUID-VkAccelerationStructureBuildGeometryInfoKHR-dstAccelerationStructure-04927"
  id="VUID-VkAccelerationStructureBuildGeometryInfoKHR-dstAccelerationStructure-04927"></a>
  VUID-VkAccelerationStructureBuildGeometryInfoKHR-dstAccelerationStructure-04927  
  If `dstAccelerationStructure` was created with
  `VK_ACCELERATION_STRUCTURE_CREATE_MOTION_BIT_NV` set in
  [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoKHR.html)::`flags`,
  `VK_BUILD_ACCELERATION_STRUCTURE_MOTION_BIT_NV` **must** be set in
  `flags`

- <a href="#VUID-VkAccelerationStructureBuildGeometryInfoKHR-flags-04928"
  id="VUID-VkAccelerationStructureBuildGeometryInfoKHR-flags-04928"></a>
  VUID-VkAccelerationStructureBuildGeometryInfoKHR-flags-04928  
  If `VK_BUILD_ACCELERATION_STRUCTURE_MOTION_BIT_NV` is set in `flags`,
  `dstAccelerationStructure` **must** have been created with
  `VK_ACCELERATION_STRUCTURE_CREATE_MOTION_BIT_NV` set in
  [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoKHR.html)::`flags`

- <a href="#VUID-VkAccelerationStructureBuildGeometryInfoKHR-flags-04929"
  id="VUID-VkAccelerationStructureBuildGeometryInfoKHR-flags-04929"></a>
  VUID-VkAccelerationStructureBuildGeometryInfoKHR-flags-04929  
  If `VK_BUILD_ACCELERATION_STRUCTURE_MOTION_BIT_NV` is set in `flags`,
  `type` **must** not be `VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR`

- <a href="#VUID-VkAccelerationStructureBuildGeometryInfoKHR-flags-07334"
  id="VUID-VkAccelerationStructureBuildGeometryInfoKHR-flags-07334"></a>
  VUID-VkAccelerationStructureBuildGeometryInfoKHR-flags-07334  
  If `flags` has the
  `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_OPACITY_MICROMAP_UPDATE_EXT`
  bit set then it **must** not have the
  `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_OPACITY_MICROMAP_DATA_UPDATE_EXT`
  bit set

Valid Usage (Implicit)

- <a href="#VUID-VkAccelerationStructureBuildGeometryInfoKHR-sType-sType"
  id="VUID-VkAccelerationStructureBuildGeometryInfoKHR-sType-sType"></a>
  VUID-VkAccelerationStructureBuildGeometryInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_BUILD_GEOMETRY_INFO_KHR`

- <a href="#VUID-VkAccelerationStructureBuildGeometryInfoKHR-pNext-pNext"
  id="VUID-VkAccelerationStructureBuildGeometryInfoKHR-pNext-pNext"></a>
  VUID-VkAccelerationStructureBuildGeometryInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a
  href="#VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-parameter"
  id="VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-parameter"></a>
  VUID-VkAccelerationStructureBuildGeometryInfoKHR-type-parameter  
  `type` **must** be a valid
  [VkAccelerationStructureTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureTypeKHR.html)
  value

- <a
  href="#VUID-VkAccelerationStructureBuildGeometryInfoKHR-flags-parameter"
  id="VUID-VkAccelerationStructureBuildGeometryInfoKHR-flags-parameter"></a>
  VUID-VkAccelerationStructureBuildGeometryInfoKHR-flags-parameter  
  `flags` **must** be a valid combination of
  [VkBuildAccelerationStructureFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildAccelerationStructureFlagBitsKHR.html)
  values

- <a
  href="#VUID-VkAccelerationStructureBuildGeometryInfoKHR-pGeometries-parameter"
  id="VUID-VkAccelerationStructureBuildGeometryInfoKHR-pGeometries-parameter"></a>
  VUID-VkAccelerationStructureBuildGeometryInfoKHR-pGeometries-parameter  
  If `geometryCount` is not `0`, and `pGeometries` is not `NULL`,
  `pGeometries` **must** be a valid pointer to an array of
  `geometryCount` valid
  [VkAccelerationStructureGeometryKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryKHR.html)
  structures

- <a
  href="#VUID-VkAccelerationStructureBuildGeometryInfoKHR-ppGeometries-parameter"
  id="VUID-VkAccelerationStructureBuildGeometryInfoKHR-ppGeometries-parameter"></a>
  VUID-VkAccelerationStructureBuildGeometryInfoKHR-ppGeometries-parameter  
  If `geometryCount` is not `0`, and `ppGeometries` is not `NULL`,
  `ppGeometries` **must** be a valid pointer to an array of
  `geometryCount` valid pointers to valid
  [VkAccelerationStructureGeometryKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryKHR.html)
  structures

- <a href="#VUID-VkAccelerationStructureBuildGeometryInfoKHR-commonparent"
  id="VUID-VkAccelerationStructureBuildGeometryInfoKHR-commonparent"></a>
  VUID-VkAccelerationStructureBuildGeometryInfoKHR-commonparent  
  Both of `dstAccelerationStructure`, and `srcAccelerationStructure`
  that are valid handles of non-ignored parameters **must** have been
  created, allocated, or retrieved from the same
  [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkAccelerationStructureGeometryKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryKHR.html),
[VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html),
[VkAccelerationStructureTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureTypeKHR.html),
[VkBuildAccelerationStructureFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildAccelerationStructureFlagsKHR.html),
[VkBuildAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildAccelerationStructureModeKHR.html),
[VkDeviceOrHostAddressKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceOrHostAddressKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkBuildAccelerationStructuresKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBuildAccelerationStructuresKHR.html),
[vkCmdBuildAccelerationStructuresIndirectKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBuildAccelerationStructuresIndirectKHR.html),
[vkCmdBuildAccelerationStructuresKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBuildAccelerationStructuresKHR.html),
[vkGetAccelerationStructureBuildSizesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureBuildSizesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAccelerationStructureBuildGeometryInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
