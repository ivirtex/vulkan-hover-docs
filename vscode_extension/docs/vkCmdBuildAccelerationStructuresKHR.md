# vkCmdBuildAccelerationStructuresKHR(3) Manual Page

## Name

vkCmdBuildAccelerationStructuresKHR - Build an acceleration structure



## [](#_c_specification)C Specification

To build acceleration structures call:

```c++
// Provided by VK_KHR_acceleration_structure
void vkCmdBuildAccelerationStructuresKHR(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    infoCount,
    const VkAccelerationStructureBuildGeometryInfoKHR* pInfos,
    const VkAccelerationStructureBuildRangeInfoKHR* const* ppBuildRangeInfos);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `infoCount` is the number of acceleration structures to build. It specifies the number of the `pInfos` structures and `ppBuildRangeInfos` pointers that **must** be provided.
- `pInfos` is a pointer to an array of `infoCount` [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html) structures defining the geometry used to build each acceleration structure.
- `ppBuildRangeInfos` is a pointer to an array of `infoCount` pointers to arrays of [VkAccelerationStructureBuildRangeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildRangeInfoKHR.html) structures. Each `ppBuildRangeInfos`\[i] is a pointer to an array of `pInfos`\[i].`geometryCount` [VkAccelerationStructureBuildRangeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildRangeInfoKHR.html) structures defining dynamic offsets to the addresses where geometry data is stored, as defined by `pInfos`\[i].

## [](#_description)Description

The `vkCmdBuildAccelerationStructuresKHR` command provides the ability to initiate multiple acceleration structures builds, however there is no ordering or synchronization implied between any of the individual acceleration structure builds.

Note

This means that an application **cannot** build a top-level acceleration structure in the same [vkCmdBuildAccelerationStructuresKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBuildAccelerationStructuresKHR.html) call as the associated bottom-level or instance acceleration structures are being built. There also **cannot** be any memory aliasing between any acceleration structure memories or scratch memories being used by any of the builds.

Accesses to the acceleration structure scratch buffers as identified by the [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)::`scratchData` buffer device addresses **must** be [synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies) with the `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) and an [access type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types) of (`VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR` | `VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_KHR`). Accesses to each [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)::`srcAccelerationStructure` and [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)::`dstAccelerationStructure` **must** be [synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies) with the `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) and an [access type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types) of `VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR` or `VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_KHR`, as appropriate.

Accesses to other input buffers as identified by any used values of [VkAccelerationStructureGeometryMotionTrianglesDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryMotionTrianglesDataNV.html)::`vertexData`, [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)::`vertexData`, [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)::`indexData`, [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)::`transformData`, [VkAccelerationStructureGeometryAabbsDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryAabbsDataKHR.html)::`data`, and [VkAccelerationStructureGeometryInstancesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryInstancesDataKHR.html)::`data` **must** be [synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies) with the `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) and an [access type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types) of `VK_ACCESS_SHADER_READ_BIT`.

Valid Usage

- [](#VUID-vkCmdBuildAccelerationStructuresKHR-accelerationStructure-08923)VUID-vkCmdBuildAccelerationStructuresKHR-accelerationStructure-08923  
  The [`VkPhysicalDeviceAccelerationStructureFeaturesKHR`::`accelerationStructure`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-accelerationStructure) feature **must** be enabled

<!--THE END-->

- [](#VUID-vkCmdBuildAccelerationStructuresKHR-mode-04628)VUID-vkCmdBuildAccelerationStructuresKHR-mode-04628  
  The `mode` member of each element of `pInfos` **must** be a valid [VkBuildAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildAccelerationStructureModeKHR.html) value
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-srcAccelerationStructure-04629)VUID-vkCmdBuildAccelerationStructuresKHR-srcAccelerationStructure-04629  
  If the `srcAccelerationStructure` member of any element of `pInfos` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the `srcAccelerationStructure` member **must** be a valid [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html) handle
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-04630)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-04630  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `srcAccelerationStructure` member **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03403)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03403  
  The `srcAccelerationStructure` member of any element of `pInfos` **must** not be the same acceleration structure as the `dstAccelerationStructure` member of any other element of `pInfos`
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03698)VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03698  
  The `dstAccelerationStructure` member of any element of `pInfos` **must** not be the same acceleration structure as the `dstAccelerationStructure` member of any other element of `pInfos`
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03800)VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03800  
  The `dstAccelerationStructure` member of any element of `pInfos` **must** be a valid [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html) handle
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03699)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03699  
  For each element of `pInfos`, if its `type` member is `VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR`, its `dstAccelerationStructure` member **must** have been created with a value of [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoKHR.html)::`type` equal to either `VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR` or `VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR`
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03700)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03700  
  For each element of `pInfos`, if its `type` member is `VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_KHR`, its `dstAccelerationStructure` member **must** have been created with a value of [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoKHR.html)::`type` equal to either `VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_KHR` or `VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR`
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03663)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03663  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, [inactive primitives](#acceleration-structure-inactive-prims) in its `srcAccelerationStructure` member **must** not be made active
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03664)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03664  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, active primitives in its `srcAccelerationStructure` member **must** not be made [inactive](#acceleration-structure-inactive-prims)
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-None-03407)VUID-vkCmdBuildAccelerationStructuresKHR-None-03407  
  The `dstAccelerationStructure` member of any element of `pInfos` **must** not be referenced by the `geometry.instances.data` member of any element of `pGeometries` or `ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR` in any other element of `pInfos`
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03701)VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03701  
  The range of memory backing the `dstAccelerationStructure` member of any element of `pInfos` that is accessed by this command **must** not overlap the memory backing the `srcAccelerationStructure` member of any other element of `pInfos` with a `mode` equal to `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, which is accessed by this command
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03702)VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03702  
  The range of memory backing the `dstAccelerationStructure` member of any element of `pInfos` that is accessed by this command **must** not overlap the memory backing the `dstAccelerationStructure` member of any other element of `pInfos`, which is accessed by this command
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03703)VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03703  
  The range of memory backing the `dstAccelerationStructure` member of any element of `pInfos` that is accessed by this command **must** not overlap the memory backing the `scratchData` member of any element of `pInfos` (including the same element), which is accessed by this command
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-scratchData-03704)VUID-vkCmdBuildAccelerationStructuresKHR-scratchData-03704  
  The range of memory backing the `scratchData` member of any element of `pInfos` that is accessed by this command **must** not overlap the memory backing the `scratchData` member of any other element of `pInfos`, which is accessed by this command
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-scratchData-03705)VUID-vkCmdBuildAccelerationStructuresKHR-scratchData-03705  
  The range of memory backing the `scratchData` member of any element of `pInfos` that is accessed by this command **must** not overlap the memory backing the `srcAccelerationStructure` member of any element of `pInfos` with a `mode` equal to `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR` (including the same element), which is accessed by this command
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03706)VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03706  
  The range of memory backing the `dstAccelerationStructure` member of any element of `pInfos` that is accessed by this command **must** not overlap the memory backing any acceleration structure referenced by the `geometry.instances.data` member of any element of `pGeometries` or `ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR` in any other element of `pInfos`, which is accessed by this command
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03667)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03667  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `srcAccelerationStructure` member **must** have previously been constructed with `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_UPDATE_BIT_KHR` set in [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)::`flags` in the build
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03668)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03668  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `srcAccelerationStructure` and `dstAccelerationStructure` members **must** either be the same [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html), or not have any [memory aliasing](#resources-memory-aliasing)
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03758)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03758  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `geometryCount` member **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03759)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03759  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `flags` member **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03760)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03760  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `type` member **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03761)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03761  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, its `geometryType` member **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03762)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03762  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, its `flags` member **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03763)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03763  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, its `geometry.triangles.vertexFormat` member **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03764)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03764  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, its `geometry.triangles.maxVertex` member **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03765)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03765  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, its `geometry.triangles.indexType` member **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03766)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03766  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if its `geometry.triangles.transformData` address was `NULL` when `srcAccelerationStructure` was last built, then it **must** be `NULL`
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03767)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03767  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if its `geometry.triangles.transformData` address was not `NULL` when `srcAccelerationStructure` was last built, then it **must** not be `NULL`
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03768)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03768  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, and `geometry.triangles.indexType` is not `VK_INDEX_TYPE_NONE_KHR`, then the value of each index referenced **must** be the same as the corresponding index value when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-primitiveCount-03769)VUID-vkCmdBuildAccelerationStructuresKHR-primitiveCount-03769  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, the `primitiveCount` member of its corresponding `VkAccelerationStructureBuildRangeInfoKHR` structure **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03801)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03801  
  For each element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR`, the corresponding `ppBuildRangeInfos`\[i]\[j].`primitiveCount` **must** be less than or equal to [VkPhysicalDeviceAccelerationStructurePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceAccelerationStructurePropertiesKHR.html)::`maxInstanceCount`

<!--THE END-->

- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03707)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03707  
  For each element of `pInfos`, the `buffer` used to create its `dstAccelerationStructure` member **must** be bound to device memory
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03708)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03708  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR` the `buffer` used to create its `srcAccelerationStructure` member **must** be bound to device memory
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03709)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03709  
  For each element of `pInfos`, the `buffer` used to create each acceleration structure referenced by the `geometry.instances.data` member of any element of `pGeometries` or `ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR` **must** be bound to device memory
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03671)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03671  
  If `pInfos`\[i].`mode` is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_BUILD_KHR`, all addresses between `pInfos`\[i].`scratchData.deviceAddress` and `pInfos`\[i].`scratchData.deviceAddress` + N - 1 **must** be in the buffer device address range of the same buffer, where N is given by the `buildScratchSize` member of the [VkAccelerationStructureBuildSizesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildSizesInfoKHR.html) structure returned from a call to [vkGetAccelerationStructureBuildSizesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureBuildSizesKHR.html) with an identical [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html) structure and primitive count
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03672)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03672  
  If `pInfos`\[i].`mode` is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, all addresses between `pInfos`\[i].`scratchData.deviceAddress` and `pInfos`\[i].`scratchData.deviceAddress` + N - 1 **must** be in the buffer device address range of the same buffer, where N is given by the `updateScratchSize` member of the [VkAccelerationStructureBuildSizesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildSizesInfoKHR.html) structure returned from a call to [vkGetAccelerationStructureBuildSizesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureBuildSizesKHR.html) with an identical [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html) structure and primitive count
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-geometry-03673)VUID-vkCmdBuildAccelerationStructuresKHR-geometry-03673  
  The buffers from which the buffer device addresses for all of the `geometry.triangles.vertexData`, `geometry.triangles.indexData`, `geometry.triangles.transformData`, `geometry.aabbs.data`, and `geometry.instances.data` members of all `pInfos`\[i].`pGeometries` and `pInfos`\[i].`ppGeometries` are queried **must** have been created with the `VK_BUFFER_USAGE_ACCELERATION_STRUCTURE_BUILD_INPUT_READ_ONLY_BIT_KHR` usage flag
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03674)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03674  
  The buffer from which the buffer device address `pInfos`\[i].`scratchData.deviceAddress` is queried **must** have been created with `VK_BUFFER_USAGE_STORAGE_BUFFER_BIT` usage flag
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03802)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03802  
  For each element of `pInfos`, its `scratchData.deviceAddress` member **must** be a valid device address obtained from [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferDeviceAddress.html)
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03803)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03803  
  For each element of `pInfos`, if `scratchData.deviceAddress` is the address of a non-sparse buffer then it **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03710)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03710  
  For each element of `pInfos`, its `scratchData.deviceAddress` member **must** be a multiple of [VkPhysicalDeviceAccelerationStructurePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceAccelerationStructurePropertiesKHR.html)::`minAccelerationStructureScratchOffsetAlignment`
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03804)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03804  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, `geometry.triangles.vertexData.deviceAddress` **must** be a valid device address obtained from [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferDeviceAddress.html)
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03805)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03805  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if `geometry.triangles.vertexData.deviceAddress` is the address of a non-sparse buffer then it **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03711)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03711  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, `geometry.triangles.vertexData.deviceAddress` **must** be aligned to the size in bytes of the smallest component of the format in `vertexFormat`
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03806)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03806  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if `geometry.triangles.indexType` is not `VK_INDEX_TYPE_NONE_KHR`, `geometry.triangles.indexData.deviceAddress` **must** be a valid device address obtained from [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferDeviceAddress.html)
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03807)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03807  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if `geometry.triangles.indexType` is not `VK_INDEX_TYPE_NONE_KHR`, if `geometry.triangles.indexData.deviceAddress` is the address of a non-sparse buffer then it **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03712)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03712  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, and with `geometry.triangles.indexType` not equal to `VK_INDEX_TYPE_NONE_KHR`, `geometry.triangles.indexData.deviceAddress` **must** be aligned to the size in bytes of the type in `indexType`
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03808)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03808  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if `geometry.triangles.transformData.deviceAddress` is not `0`, it **must** be a valid device address obtained from [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferDeviceAddress.html)
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03809)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03809  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if `geometry.triangles.transformData.deviceAddress` is the address of a non-sparse buffer then it **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03810)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03810  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if `geometry.triangles.transformData.deviceAddress` is not `0`, it **must** be aligned to `16` bytes
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03811)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03811  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_AABBS_KHR`, `geometry.aabbs.data.deviceAddress` **must** be a valid device address obtained from [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferDeviceAddress.html)
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03812)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03812  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_AABBS_KHR`, if `geometry.aabbs.data.deviceAddress` is the address of a non-sparse buffer then it **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03714)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03714  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_AABBS_KHR`, `geometry.aabbs.data.deviceAddress` **must** be aligned to `8` bytes
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03715)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03715  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR`, if `geometry.arrayOfPointers` is `VK_FALSE`, `geometry.instances.data.deviceAddress` **must** be aligned to `16` bytes
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03716)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03716  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR`, if `geometry.arrayOfPointers` is `VK_TRUE`, `geometry.instances.data.deviceAddress` **must** be aligned to `8` bytes
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03717)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03717  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR`, if `geometry.arrayOfPointers` is `VK_TRUE`, each element of `geometry.instances.data.deviceAddress` in device memory **must** be aligned to `16` bytes
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03813)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03813  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR`, `geometry.instances.data.deviceAddress` **must** be a valid device address obtained from [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferDeviceAddress.html)
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03814)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03814  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR`, if `geometry.instances.data.deviceAddress` is the address of a non-sparse buffer then it **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-06707)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-06707  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR`, each [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInstanceKHR.html)::`accelerationStructureReference` value in `geometry.instances.data.deviceAddress` **must** be a valid device address containing a value obtained from [vkGetAccelerationStructureDeviceAddressKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureDeviceAddressKHR.html) or `0`
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-10607)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-10607  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR`, if `VK_GEOMETRY_INSTANCE_DISABLE_OPACITY_MICROMAPS_BIT_EXT` is set in [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInstanceKHR.html)::`flags` then `geometry.instances.data.deviceAddress` **must** refer to an acceleration structure that was built with `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_DISABLE_OPACITY_MICROMAPS_BIT_EXT` set in [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)::`flags`
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-commandBuffer-09547)VUID-vkCmdBuildAccelerationStructuresKHR-commandBuffer-09547  
  `commandBuffer` **must** not be a protected command buffer

<!--THE END-->

- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-10126)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-10126  
  For each `pInfos`\[i], `dstAccelerationStructure` **must** have been created with a value of [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoKHR.html)::`size` greater than or equal to either:
  
  - the memory size required by the build operation, as returned by [vkGetAccelerationStructureBuildSizesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureBuildSizesKHR.html) with `pBuildInfo` = `pInfos`\[i] and with each element of the `pMaxPrimitiveCounts` array greater than or equal to the equivalent `ppBuildRangeInfos`\[i]\[j].`primitiveCount` values for `j` in \[0,`pInfos`\[i].`geometryCount`) or,
  - the result of querying the corresponding `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_KHR`, if updating a compacted acceleration structure
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-ppBuildRangeInfos-03676)VUID-vkCmdBuildAccelerationStructuresKHR-ppBuildRangeInfos-03676  
  Each element of `ppBuildRangeInfos`\[i] **must** be a valid pointer to an array of `pInfos`\[i].`geometryCount` `VkAccelerationStructureBuildRangeInfoKHR` structures

Valid Usage (Implicit)

- [](#VUID-vkCmdBuildAccelerationStructuresKHR-commandBuffer-parameter)VUID-vkCmdBuildAccelerationStructuresKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-parameter)VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-parameter  
  `pInfos` **must** be a valid pointer to an array of `infoCount` valid [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html) structures
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-ppBuildRangeInfos-parameter)VUID-vkCmdBuildAccelerationStructuresKHR-ppBuildRangeInfos-parameter  
  `ppBuildRangeInfos` **must** be a valid pointer to an array of `infoCount` [VkAccelerationStructureBuildRangeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildRangeInfoKHR.html) structures
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-commandBuffer-recording)VUID-vkCmdBuildAccelerationStructuresKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-commandBuffer-cmdpool)VUID-vkCmdBuildAccelerationStructuresKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support compute operations
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-renderpass)VUID-vkCmdBuildAccelerationStructuresKHR-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-videocoding)VUID-vkCmdBuildAccelerationStructuresKHR-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdBuildAccelerationStructuresKHR-infoCount-arraylength)VUID-vkCmdBuildAccelerationStructuresKHR-infoCount-arraylength  
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

Compute

Action

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html), [VkAccelerationStructureBuildRangeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildRangeInfoKHR.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBuildAccelerationStructuresKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0