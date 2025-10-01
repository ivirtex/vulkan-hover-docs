# vkCmdBuildAccelerationStructuresIndirectKHR(3) Manual Page

## Name

vkCmdBuildAccelerationStructuresIndirectKHR - Build an acceleration structure with some parameters provided on the device



## [](#_c_specification)C Specification

To build acceleration structures with some parameters sourced on the device call:

```c++
// Provided by VK_KHR_acceleration_structure
void vkCmdBuildAccelerationStructuresIndirectKHR(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    infoCount,
    const VkAccelerationStructureBuildGeometryInfoKHR* pInfos,
    const VkDeviceAddress*                      pIndirectDeviceAddresses,
    const uint32_t*                             pIndirectStrides,
    const uint32_t* const*                      ppMaxPrimitiveCounts);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `infoCount` is the number of acceleration structures to build.
- `pInfos` is a pointer to an array of `infoCount` [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html) structures defining the geometry used to build each acceleration structure.
- `pIndirectDeviceAddresses` is a pointer to an array of `infoCount` buffer device addresses which point to `pInfos`\[i].`geometryCount` [VkAccelerationStructureBuildRangeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildRangeInfoKHR.html) structures defining dynamic offsets to the addresses where geometry data is stored, as defined by `pInfos`\[i].
- `pIndirectStrides` is a pointer to an array of `infoCount` byte strides between elements of `pIndirectDeviceAddresses`.
- `ppMaxPrimitiveCounts` is a pointer to an array of `infoCount` pointers to arrays of `pInfos`\[i].`geometryCount` values indicating the maximum number of primitives that will be built by this command for each geometry.

## [](#_description)Description

Accesses to acceleration structures, scratch buffers, vertex buffers, index buffers, and instance buffers **must** be synchronized as with [vkCmdBuildAccelerationStructuresKHR](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#acceleration-structure-scratch).

Accesses to any element of `pIndirectDeviceAddresses` **must** be [synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies) with the `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) and an [access type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types) of `VK_ACCESS_INDIRECT_COMMAND_READ_BIT`.

Valid Usage

- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-accelerationStructureIndirectBuild-03650)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-accelerationStructureIndirectBuild-03650  
  The [`VkPhysicalDeviceAccelerationStructureFeaturesKHR`::`accelerationStructureIndirectBuild`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-accelerationStructureIndirectBuild) feature **must** be enabled

<!--THE END-->

- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-mode-04628)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-mode-04628  
  The `mode` member of each element of `pInfos` **must** be a valid [VkBuildAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildAccelerationStructureModeKHR.html) value
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-srcAccelerationStructure-04629)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-srcAccelerationStructure-04629  
  If the `srcAccelerationStructure` member of any element of `pInfos` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the `srcAccelerationStructure` member **must** be a valid [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html) handle
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-04630)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-04630  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `srcAccelerationStructure` member **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03403)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03403  
  The `srcAccelerationStructure` member of any element of `pInfos` **must** not be the same acceleration structure as the `dstAccelerationStructure` member of any other element of `pInfos`
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03698)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03698  
  The `dstAccelerationStructure` member of any element of `pInfos` **must** not be the same acceleration structure as the `dstAccelerationStructure` member of any other element of `pInfos`
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03800)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03800  
  The `dstAccelerationStructure` member of any element of `pInfos` **must** be a valid [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html) handle
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03699)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03699  
  For each element of `pInfos`, if its `type` member is `VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR`, its `dstAccelerationStructure` member **must** have been created with a value of [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoKHR.html)::`type` equal to either `VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR` or `VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR`
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03700)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03700  
  For each element of `pInfos`, if its `type` member is `VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_KHR`, its `dstAccelerationStructure` member **must** have been created with a value of [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoKHR.html)::`type` equal to either `VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_KHR` or `VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR`
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03663)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03663  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, [inactive primitives](#acceleration-structure-inactive-prims) in its `srcAccelerationStructure` member **must** not be made active
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03664)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03664  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, active primitives in its `srcAccelerationStructure` member **must** not be made [inactive](#acceleration-structure-inactive-prims)
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-None-03407)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-None-03407  
  The `dstAccelerationStructure` member of any element of `pInfos` **must** not be referenced by the `geometry.instances.data` member of any element of `pGeometries` or `ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR` in any other element of `pInfos`
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03701)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03701  
  The range of memory backing the `dstAccelerationStructure` member of any element of `pInfos` that is accessed by this command **must** not overlap the memory backing the `srcAccelerationStructure` member of any other element of `pInfos` with a `mode` equal to `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, which is accessed by this command
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03702)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03702  
  The range of memory backing the `dstAccelerationStructure` member of any element of `pInfos` that is accessed by this command **must** not overlap the memory backing the `dstAccelerationStructure` member of any other element of `pInfos`, which is accessed by this command
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03703)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03703  
  The range of memory backing the `dstAccelerationStructure` member of any element of `pInfos` that is accessed by this command **must** not overlap the memory backing the `scratchData` member of any element of `pInfos` (including the same element), which is accessed by this command
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-scratchData-03704)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-scratchData-03704  
  The range of memory backing the `scratchData` member of any element of `pInfos` that is accessed by this command **must** not overlap the memory backing the `scratchData` member of any other element of `pInfos`, which is accessed by this command
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-scratchData-03705)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-scratchData-03705  
  The range of memory backing the `scratchData` member of any element of `pInfos` that is accessed by this command **must** not overlap the memory backing the `srcAccelerationStructure` member of any element of `pInfos` with a `mode` equal to `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR` (including the same element), which is accessed by this command
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03706)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03706  
  The range of memory backing the `dstAccelerationStructure` member of any element of `pInfos` that is accessed by this command **must** not overlap the memory backing any acceleration structure referenced by the `geometry.instances.data` member of any element of `pGeometries` or `ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR` in any other element of `pInfos`, which is accessed by this command
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03667)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03667  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `srcAccelerationStructure` member **must** have previously been constructed with `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_UPDATE_BIT_KHR` set in [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)::`flags` in the build
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03668)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03668  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `srcAccelerationStructure` and `dstAccelerationStructure` members **must** either be the same [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html), or not have any [memory aliasing](#resources-memory-aliasing)
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03758)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03758  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `geometryCount` member **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03759)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03759  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `flags` member **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03760)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03760  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `type` member **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03761)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03761  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, its `geometryType` member **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03762)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03762  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, its `flags` member **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03763)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03763  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, its `geometry.triangles.vertexFormat` member **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03764)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03764  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, its `geometry.triangles.maxVertex` member **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03765)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03765  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, its `geometry.triangles.indexType` member **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03766)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03766  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if its `geometry.triangles.transformData` address was `NULL` when `srcAccelerationStructure` was last built, then it **must** be `NULL`
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03767)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03767  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if its `geometry.triangles.transformData` address was not `NULL` when `srcAccelerationStructure` was last built, then it **must** not be `NULL`
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-10898)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-10898  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_DENSE_GEOMETRY_FORMAT_TRIANGLES_AMDX`, the `numTriangles` member of the `VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX` structure in the `pNext` chain **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-10899)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-10899  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_DENSE_GEOMETRY_FORMAT_TRIANGLES_AMDX`, the `numVertices` member of the `VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX` structure in the `pNext` chain **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-10900)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-10900  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_DENSE_GEOMETRY_FORMAT_TRIANGLES_AMDX`, the `maxPrimitiveIndex` member of the `VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX` structure in the `pNext` chain **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-10901)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-10901  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_DENSE_GEOMETRY_FORMAT_TRIANGLES_AMDX`, the `maxGeometryIndex` member of the `VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX` structure in the `pNext` chain **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-10902)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-10902  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_DENSE_GEOMETRY_FORMAT_TRIANGLES_AMDX`, the `format` member of the `VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX` structure in the `pNext` chain **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-10903)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-10903  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_DENSE_GEOMETRY_FORMAT_TRIANGLES_AMDX`, the `dataSize` member of the `VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX` structure in the `pNext` chain **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03768)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03768  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, and `geometry.triangles.indexType` is not `VK_INDEX_TYPE_NONE_KHR`, then the value of each index referenced **must** be the same as the corresponding index value when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-primitiveCount-03769)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-primitiveCount-03769  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, the `primitiveCount` member of its corresponding `VkAccelerationStructureBuildRangeInfoKHR` structure **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03801)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03801  
  For each element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR`, the corresponding `ppMaxPrimitiveCounts`\[i]\[j] **must** be less than or equal to [VkPhysicalDeviceAccelerationStructurePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceAccelerationStructurePropertiesKHR.html)::`maxInstanceCount`

<!--THE END-->

- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03707)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03707  
  For each element of `pInfos`, the `buffer` used to create its `dstAccelerationStructure` member **must** be bound to device memory
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03708)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03708  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR` the `buffer` used to create its `srcAccelerationStructure` member **must** be bound to device memory
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03709)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03709  
  For each element of `pInfos`, the `buffer` used to create each acceleration structure referenced by the `geometry.instances.data` member of any element of `pGeometries` or `ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR` **must** be bound to device memory
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03671)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03671  
  If `pInfos`\[i].`mode` is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_BUILD_KHR`, all addresses between `pInfos`\[i].`scratchData.deviceAddress` and `pInfos`\[i].`scratchData.deviceAddress` + N - 1 **must** be in the buffer device address range of the same buffer, where N is given by the `buildScratchSize` member of the [VkAccelerationStructureBuildSizesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildSizesInfoKHR.html) structure returned from a call to [vkGetAccelerationStructureBuildSizesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureBuildSizesKHR.html) with an identical [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html) structure and primitive count
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03672)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03672  
  If `pInfos`\[i].`mode` is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, all addresses between `pInfos`\[i].`scratchData.deviceAddress` and `pInfos`\[i].`scratchData.deviceAddress` + N - 1 **must** be in the buffer device address range of the same buffer, where N is given by the `updateScratchSize` member of the [VkAccelerationStructureBuildSizesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildSizesInfoKHR.html) structure returned from a call to [vkGetAccelerationStructureBuildSizesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureBuildSizesKHR.html) with an identical [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html) structure and primitive count
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-geometry-03673)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-geometry-03673  
  The buffers from which the buffer device addresses for all of the `geometry.triangles.vertexData`, `geometry.triangles.indexData`, `geometry.triangles.transformData`, `geometry.aabbs.data`, and `geometry.instances.data` members of all `pInfos`\[i].`pGeometries` and `pInfos`\[i].`ppGeometries` are queried **must** have been created with the `VK_BUFFER_USAGE_ACCELERATION_STRUCTURE_BUILD_INPUT_READ_ONLY_BIT_KHR` usage flag
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03674)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03674  
  The buffer from which the buffer device address `pInfos`\[i].`scratchData.deviceAddress` is queried **must** have been created with `VK_BUFFER_USAGE_STORAGE_BUFFER_BIT` usage flag
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03802)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03802  
  For each element of `pInfos`, its `scratchData.deviceAddress` member **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03710)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03710  
  For each element of `pInfos`, its `scratchData.deviceAddress` member **must** be a multiple of [VkPhysicalDeviceAccelerationStructurePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceAccelerationStructurePropertiesKHR.html)::`minAccelerationStructureScratchOffsetAlignment`
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03804)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03804  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, `geometry.triangles.vertexData.deviceAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03711)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03711  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, `geometry.triangles.vertexData.deviceAddress` **must** be aligned to the size in bytes of the smallest component of the format in `vertexFormat`
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03806)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03806  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if `geometry.triangles.indexType` is not `VK_INDEX_TYPE_NONE_KHR`, `geometry.triangles.indexData.deviceAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03712)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03712  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, and with `geometry.triangles.indexType` not equal to `VK_INDEX_TYPE_NONE_KHR`, `geometry.triangles.indexData.deviceAddress` **must** be aligned to the size in bytes of the type in `indexType`
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03808)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03808  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if `geometry.triangles.transformData.deviceAddress` is not `0`, it **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03810)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03810  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if `geometry.triangles.transformData.deviceAddress` is not `0`, it **must** be aligned to `16` bytes
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03811)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03811  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_AABBS_KHR`, `geometry.aabbs.data.deviceAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03714)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03714  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_AABBS_KHR`, `geometry.aabbs.data.deviceAddress` **must** be aligned to `8` bytes
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03715)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03715  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR`, if `geometry.arrayOfPointers` is `VK_FALSE`, `geometry.instances.data.deviceAddress` **must** be aligned to `16` bytes
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03716)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03716  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR`, if `geometry.arrayOfPointers` is `VK_TRUE`, `geometry.instances.data.deviceAddress` **must** be aligned to `8` bytes
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03717)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03717  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR`, if `geometry.arrayOfPointers` is `VK_TRUE`, each element of `geometry.instances.data.deviceAddress` in device memory **must** be aligned to `16` bytes
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03813)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03813  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR`, `geometry.instances.data.deviceAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-06707)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-06707  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR`, each [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInstanceKHR.html)::`accelerationStructureReference` value in `geometry.instances.data.deviceAddress` **must** be `0` or a value obtained from [vkGetAccelerationStructureDeviceAddressKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureDeviceAddressKHR.html) for a valid acceleration structure
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-10607)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-10607  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR`, if `VK_GEOMETRY_INSTANCE_DISABLE_OPACITY_MICROMAPS_BIT_EXT` is set in [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInstanceKHR.html)::`flags` then `geometry.instances.data.deviceAddress` **must** refer to an acceleration structure that was built with `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_DISABLE_OPACITY_MICROMAPS_BIT_EXT` set in [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)::`flags`
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-commandBuffer-09547)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-commandBuffer-09547  
  `commandBuffer` **must** not be a protected command buffer
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-10904)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-10904  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if there is an instance of [VkAccelerationStructureTrianglesOpacityMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureTrianglesOpacityMicromapEXT.html) in the `geometry.triangles.pNext` chain, and its `indexType` is `VK_INDEX_TYPE_NONE_KHR`, then its `indexBuffer.deviceAddress` **must** be 0
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-10905)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-10905  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if there is an instance of [VkAccelerationStructureTrianglesOpacityMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureTrianglesOpacityMicromapEXT.html) in the `geometry.triangles.pNext` chain, and its `indexType` is not `VK_INDEX_TYPE_NONE_KHR`, then its `indexBuffer.deviceAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-03646)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-03646  
  For any element of `pIndirectDeviceAddresses`\[i], all device addresses between `pIndirectDeviceAddresses`\[i] and `pIndirectDeviceAddresses`\[i] + (`pInfos`\[i].`geometryCount` × `pIndirectStrides`\[i]) - 1 **must** be in the buffer device address range of the same buffer
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-03647)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-03647  
  For any element of `pIndirectDeviceAddresses`, the buffer from which it was queried **must** have been created with the `VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT` bit set
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-03648)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-03648  
  Each element of `pIndirectDeviceAddresses` **must** be a multiple of `4`
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectStrides-03787)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectStrides-03787  
  Each element of `pIndirectStrides` **must** be a multiple of `4`
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-03651)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-03651  
  Each [VkAccelerationStructureBuildRangeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildRangeInfoKHR.html) structure referenced by any element of `pIndirectDeviceAddresses` **must** be a valid [VkAccelerationStructureBuildRangeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildRangeInfoKHR.html) structure
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03652)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03652  
  `pInfos`\[i].`dstAccelerationStructure` **must** have been created with a value of [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoKHR.html)::`size` greater than or equal to the memory size required by the build operation, as returned by [vkGetAccelerationStructureBuildSizesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureBuildSizesKHR.html) with `pBuildInfo` = `pInfos`\[i] and `pMaxPrimitiveCounts` = `ppMaxPrimitiveCounts`\[i]
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-ppMaxPrimitiveCounts-03653)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-ppMaxPrimitiveCounts-03653  
  Each `ppMaxPrimitiveCounts`\[i]\[j] **must** be greater than or equal to the `primitiveCount` value specified by the [VkAccelerationStructureBuildRangeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildRangeInfoKHR.html) structure located at `pIndirectDeviceAddresses`\[i] + (`j` × `pIndirectStrides`\[i])

Valid Usage (Implicit)

- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-commandBuffer-parameter)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-parameter)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-parameter  
  `pInfos` **must** be a valid pointer to an array of `infoCount` valid [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html) structures
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-parameter)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-parameter  
  `pIndirectDeviceAddresses` **must** be a valid pointer to an array of `infoCount` [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) values
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectStrides-parameter)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectStrides-parameter  
  `pIndirectStrides` **must** be a valid pointer to an array of `infoCount` `uint32_t` values
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-ppMaxPrimitiveCounts-parameter)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-ppMaxPrimitiveCounts-parameter  
  `ppMaxPrimitiveCounts` **must** be a valid pointer to an array of `infoCount` `uint32_t` values
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-commandBuffer-recording)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-commandBuffer-cmdpool)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_COMPUTE\_BIT operations
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-renderpass)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-videocoding)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-infoCount-arraylength)VUID-vkCmdBuildAccelerationStructuresIndirectKHR-infoCount-arraylength  
  `infoCount` **must** be greater than `0`

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Outside

Outside

VK\_QUEUE\_COMPUTE\_BIT

Action

Conditional Rendering

vkCmdBuildAccelerationStructuresIndirectKHR is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBuildAccelerationStructuresIndirectKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0