# vkBuildAccelerationStructuresKHR(3) Manual Page

## Name

vkBuildAccelerationStructuresKHR - Build an acceleration structure on the host



## [](#_c_specification)C Specification

To build acceleration structures on the host, call:

```c++
// Provided by VK_KHR_acceleration_structure
VkResult vkBuildAccelerationStructuresKHR(
    VkDevice                                    device,
    VkDeferredOperationKHR                      deferredOperation,
    uint32_t                                    infoCount,
    const VkAccelerationStructureBuildGeometryInfoKHR* pInfos,
    const VkAccelerationStructureBuildRangeInfoKHR* const* ppBuildRangeInfos);
```

## [](#_parameters)Parameters

- `device` is the `VkDevice` for which the acceleration structures are being built.
- `deferredOperation` is an optional [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html) to [request deferral](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#deferred-host-operations-requesting) for this command.
- `infoCount` is the number of acceleration structures to build. It specifies the number of the `pInfos` structures and `ppBuildRangeInfos` pointers that **must** be provided.
- `pInfos` is a pointer to an array of `infoCount` [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html) structures defining the geometry used to build each acceleration structure.
- `ppBuildRangeInfos` is a pointer to an array of `infoCount` pointers to arrays of [VkAccelerationStructureBuildRangeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildRangeInfoKHR.html) structures. Each `ppBuildRangeInfos`\[i] is a pointer to an array of `pInfos`\[i].`geometryCount` [VkAccelerationStructureBuildRangeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildRangeInfoKHR.html) structures defining dynamic offsets to the addresses where geometry data is stored, as defined by `pInfos`\[i].

## [](#_description)Description

This command fulfills the same task as [vkCmdBuildAccelerationStructuresKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBuildAccelerationStructuresKHR.html) but is executed by the host.

The `vkBuildAccelerationStructuresKHR` command provides the ability to initiate multiple acceleration structures builds, however there is no ordering or synchronization implied between any of the individual acceleration structure builds.

Note

This means that an application **cannot** build a top-level acceleration structure in the same [vkBuildAccelerationStructuresKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBuildAccelerationStructuresKHR.html) call as the associated bottom-level or instance acceleration structures are being built. There also **cannot** be any memory aliasing between any acceleration structure memories or scratch memories being used by any of the builds.

Valid Usage

- [](#VUID-vkBuildAccelerationStructuresKHR-accelerationStructureHostCommands-03581)VUID-vkBuildAccelerationStructuresKHR-accelerationStructureHostCommands-03581  
  The [`VkPhysicalDeviceAccelerationStructureFeaturesKHR`::`accelerationStructureHostCommands`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-accelerationStructureHostCommands) feature **must** be enabled

<!--THE END-->

- [](#VUID-vkBuildAccelerationStructuresKHR-mode-04628)VUID-vkBuildAccelerationStructuresKHR-mode-04628  
  The `mode` member of each element of `pInfos` **must** be a valid [VkBuildAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildAccelerationStructureModeKHR.html) value
- [](#VUID-vkBuildAccelerationStructuresKHR-srcAccelerationStructure-04629)VUID-vkBuildAccelerationStructuresKHR-srcAccelerationStructure-04629  
  If the `srcAccelerationStructure` member of any element of `pInfos` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the `srcAccelerationStructure` member **must** be a valid [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html) handle
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-04630)VUID-vkBuildAccelerationStructuresKHR-pInfos-04630  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `srcAccelerationStructure` member **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03403)VUID-vkBuildAccelerationStructuresKHR-pInfos-03403  
  The `srcAccelerationStructure` member of any element of `pInfos` **must** not be the same acceleration structure as the `dstAccelerationStructure` member of any other element of `pInfos`
- [](#VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03698)VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03698  
  The `dstAccelerationStructure` member of any element of `pInfos` **must** not be the same acceleration structure as the `dstAccelerationStructure` member of any other element of `pInfos`
- [](#VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03800)VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03800  
  The `dstAccelerationStructure` member of any element of `pInfos` **must** be a valid [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html) handle
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03699)VUID-vkBuildAccelerationStructuresKHR-pInfos-03699  
  For each element of `pInfos`, if its `type` member is `VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR`, its `dstAccelerationStructure` member **must** have been created with a value of [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoKHR.html)::`type` equal to either `VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR` or `VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR`
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03700)VUID-vkBuildAccelerationStructuresKHR-pInfos-03700  
  For each element of `pInfos`, if its `type` member is `VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_KHR`, its `dstAccelerationStructure` member **must** have been created with a value of [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoKHR.html)::`type` equal to either `VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_KHR` or `VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR`
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03663)VUID-vkBuildAccelerationStructuresKHR-pInfos-03663  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, [inactive primitives](#acceleration-structure-inactive-prims) in its `srcAccelerationStructure` member **must** not be made active
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03664)VUID-vkBuildAccelerationStructuresKHR-pInfos-03664  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, active primitives in its `srcAccelerationStructure` member **must** not be made [inactive](#acceleration-structure-inactive-prims)
- [](#VUID-vkBuildAccelerationStructuresKHR-None-03407)VUID-vkBuildAccelerationStructuresKHR-None-03407  
  The `dstAccelerationStructure` member of any element of `pInfos` **must** not be referenced by the `geometry.instances.data` member of any element of `pGeometries` or `ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR` in any other element of `pInfos`
- [](#VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03701)VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03701  
  The range of memory backing the `dstAccelerationStructure` member of any element of `pInfos` that is accessed by this command **must** not overlap the memory backing the `srcAccelerationStructure` member of any other element of `pInfos` with a `mode` equal to `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, which is accessed by this command
- [](#VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03702)VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03702  
  The range of memory backing the `dstAccelerationStructure` member of any element of `pInfos` that is accessed by this command **must** not overlap the memory backing the `dstAccelerationStructure` member of any other element of `pInfos`, which is accessed by this command
- [](#VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03703)VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03703  
  The range of memory backing the `dstAccelerationStructure` member of any element of `pInfos` that is accessed by this command **must** not overlap the memory backing the `scratchData` member of any element of `pInfos` (including the same element), which is accessed by this command
- [](#VUID-vkBuildAccelerationStructuresKHR-scratchData-03704)VUID-vkBuildAccelerationStructuresKHR-scratchData-03704  
  The range of memory backing the `scratchData` member of any element of `pInfos` that is accessed by this command **must** not overlap the memory backing the `scratchData` member of any other element of `pInfos`, which is accessed by this command
- [](#VUID-vkBuildAccelerationStructuresKHR-scratchData-03705)VUID-vkBuildAccelerationStructuresKHR-scratchData-03705  
  The range of memory backing the `scratchData` member of any element of `pInfos` that is accessed by this command **must** not overlap the memory backing the `srcAccelerationStructure` member of any element of `pInfos` with a `mode` equal to `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR` (including the same element), which is accessed by this command
- [](#VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03706)VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03706  
  The range of memory backing the `dstAccelerationStructure` member of any element of `pInfos` that is accessed by this command **must** not overlap the memory backing any acceleration structure referenced by the `geometry.instances.data` member of any element of `pGeometries` or `ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR` in any other element of `pInfos`, which is accessed by this command
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03667)VUID-vkBuildAccelerationStructuresKHR-pInfos-03667  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `srcAccelerationStructure` member **must** have previously been constructed with `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_UPDATE_BIT_KHR` set in [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)::`flags` in the build
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03668)VUID-vkBuildAccelerationStructuresKHR-pInfos-03668  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `srcAccelerationStructure` and `dstAccelerationStructure` members **must** either be the same [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html), or not have any [memory aliasing](#resources-memory-aliasing)
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03758)VUID-vkBuildAccelerationStructuresKHR-pInfos-03758  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `geometryCount` member **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03759)VUID-vkBuildAccelerationStructuresKHR-pInfos-03759  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `flags` member **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03760)VUID-vkBuildAccelerationStructuresKHR-pInfos-03760  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `type` member **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03761)VUID-vkBuildAccelerationStructuresKHR-pInfos-03761  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, its `geometryType` member **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03762)VUID-vkBuildAccelerationStructuresKHR-pInfos-03762  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, its `flags` member **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03763)VUID-vkBuildAccelerationStructuresKHR-pInfos-03763  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, its `geometry.triangles.vertexFormat` member **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03764)VUID-vkBuildAccelerationStructuresKHR-pInfos-03764  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, its `geometry.triangles.maxVertex` member **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03765)VUID-vkBuildAccelerationStructuresKHR-pInfos-03765  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, its `geometry.triangles.indexType` member **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03766)VUID-vkBuildAccelerationStructuresKHR-pInfos-03766  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if its `geometry.triangles.transformData` address was `NULL` when `srcAccelerationStructure` was last built, then it **must** be `NULL`
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03767)VUID-vkBuildAccelerationStructuresKHR-pInfos-03767  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if its `geometry.triangles.transformData` address was not `NULL` when `srcAccelerationStructure` was last built, then it **must** not be `NULL`
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-10898)VUID-vkBuildAccelerationStructuresKHR-pInfos-10898  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_DENSE_GEOMETRY_FORMAT_TRIANGLES_AMDX`, the `numTriangles` member of the `VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX` structure in the `pNext` chain **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-10899)VUID-vkBuildAccelerationStructuresKHR-pInfos-10899  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_DENSE_GEOMETRY_FORMAT_TRIANGLES_AMDX`, the `numVertices` member of the `VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX` structure in the `pNext` chain **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-10900)VUID-vkBuildAccelerationStructuresKHR-pInfos-10900  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_DENSE_GEOMETRY_FORMAT_TRIANGLES_AMDX`, the `maxPrimitiveIndex` member of the `VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX` structure in the `pNext` chain **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-10901)VUID-vkBuildAccelerationStructuresKHR-pInfos-10901  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_DENSE_GEOMETRY_FORMAT_TRIANGLES_AMDX`, the `maxGeometryIndex` member of the `VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX` structure in the `pNext` chain **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-10902)VUID-vkBuildAccelerationStructuresKHR-pInfos-10902  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_DENSE_GEOMETRY_FORMAT_TRIANGLES_AMDX`, the `format` member of the `VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX` structure in the `pNext` chain **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-10903)VUID-vkBuildAccelerationStructuresKHR-pInfos-10903  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_DENSE_GEOMETRY_FORMAT_TRIANGLES_AMDX`, the `dataSize` member of the `VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX` structure in the `pNext` chain **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03768)VUID-vkBuildAccelerationStructuresKHR-pInfos-03768  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, if `geometryType` is `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, and `geometry.triangles.indexType` is not `VK_INDEX_TYPE_NONE_KHR`, then the value of each index referenced **must** be the same as the corresponding index value when `srcAccelerationStructure` was last built
- [](#VUID-vkBuildAccelerationStructuresKHR-primitiveCount-03769)VUID-vkBuildAccelerationStructuresKHR-primitiveCount-03769  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each `VkAccelerationStructureGeometryKHR` structure referred to by its `pGeometries` or `ppGeometries` members, the `primitiveCount` member of its corresponding `VkAccelerationStructureBuildRangeInfoKHR` structure **must** have the same value which was specified when `srcAccelerationStructure` was last built
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03801)VUID-vkBuildAccelerationStructuresKHR-pInfos-03801  
  For each element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR`, the corresponding `ppBuildRangeInfos`\[i]\[j].`primitiveCount` **must** be less than or equal to [VkPhysicalDeviceAccelerationStructurePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceAccelerationStructurePropertiesKHR.html)::`maxInstanceCount`

<!--THE END-->

- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-10126)VUID-vkBuildAccelerationStructuresKHR-pInfos-10126  
  For each `pInfos`\[i], `dstAccelerationStructure` **must** have been created with a value of [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoKHR.html)::`size` greater than or equal to either:
  
  - the memory size required by the build operation, as returned by [vkGetAccelerationStructureBuildSizesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureBuildSizesKHR.html) with `pBuildInfo` = `pInfos`\[i] and with each element of the `pMaxPrimitiveCounts` array greater than or equal to the equivalent `ppBuildRangeInfos`\[i]\[j].`primitiveCount` values for `j` in \[0,`pInfos`\[i].`geometryCount`) or,
  - the result of querying the corresponding `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_KHR`, if updating a compacted acceleration structure
- [](#VUID-vkBuildAccelerationStructuresKHR-ppBuildRangeInfos-03676)VUID-vkBuildAccelerationStructuresKHR-ppBuildRangeInfos-03676  
  Each element of `ppBuildRangeInfos`\[i] **must** be a valid pointer to an array of `pInfos`\[i].`geometryCount` `VkAccelerationStructureBuildRangeInfoKHR` structures , or `NULL`
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-10906)VUID-vkBuildAccelerationStructuresKHR-pInfos-10906  
  For each element of `pInfos`\[i] whose `pGeometries` or `ppGeometries` members have a `geometryType` of `VK_GEOMETRY_TYPE_DENSE_GEOMETRY_FORMAT_TRIANGLES_AMDX`, `ppBuildRangeInfos`\[i] **must** be `NULL`

<!--THE END-->

- [](#VUID-vkBuildAccelerationStructuresKHR-deferredOperation-03678)VUID-vkBuildAccelerationStructuresKHR-deferredOperation-03678  
  Any previous deferred operation that was associated with `deferredOperation` **must** be complete
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03722)VUID-vkBuildAccelerationStructuresKHR-pInfos-03722  
  For each element of `pInfos`, the `buffer` used to create its `dstAccelerationStructure` member **must** be bound to host-visible device memory
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03723)VUID-vkBuildAccelerationStructuresKHR-pInfos-03723  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR` the `buffer` used to create its `srcAccelerationStructure` member **must** be bound to host-visible device memory
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03724)VUID-vkBuildAccelerationStructuresKHR-pInfos-03724  
  For each element of `pInfos`, the `buffer` used to create each acceleration structure referenced by the `geometry.instances.data` member of any element of `pGeometries` or `ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR` **must** be bound to host-visible device memory
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03725)VUID-vkBuildAccelerationStructuresKHR-pInfos-03725  
  If `pInfos`\[i].`mode` is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_BUILD_KHR`, all addresses between `pInfos`\[i].`scratchData.hostAddress` and `pInfos`\[i].`scratchData.hostAddress` + N - 1 **must** be valid host memory, where N is given by the `buildScratchSize` member of the [VkAccelerationStructureBuildSizesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildSizesInfoKHR.html) structure returned from a call to [vkGetAccelerationStructureBuildSizesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureBuildSizesKHR.html) with an identical [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html) structure and primitive count
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03726)VUID-vkBuildAccelerationStructuresKHR-pInfos-03726  
  If `pInfos`\[i].`mode` is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, all addresses between `pInfos`\[i].`scratchData.hostAddress` and `pInfos`\[i].`scratchData.hostAddress` + N - 1 **must** be valid host memory, where N is given by the `updateScratchSize` member of the [VkAccelerationStructureBuildSizesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildSizesInfoKHR.html) structure returned from a call to [vkGetAccelerationStructureBuildSizesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureBuildSizesKHR.html) with an identical [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html) structure and primitive count
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03771)VUID-vkBuildAccelerationStructuresKHR-pInfos-03771  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, `geometry.triangles.vertexData.hostAddress` **must** be a valid host address
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03772)VUID-vkBuildAccelerationStructuresKHR-pInfos-03772  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if `geometry.triangles.indexType` is not `VK_INDEX_TYPE_NONE_KHR`, `geometry.triangles.indexData.hostAddress` **must** be a valid host address
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03773)VUID-vkBuildAccelerationStructuresKHR-pInfos-03773  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if `geometry.triangles.transformData.hostAddress` is not `0`, it **must** be a valid host address
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03774)VUID-vkBuildAccelerationStructuresKHR-pInfos-03774  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_AABBS_KHR`, `geometry.aabbs.data.hostAddress` **must** be a valid host address
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03775)VUID-vkBuildAccelerationStructuresKHR-pInfos-03775  
  For each element of `pInfos`, the `buffer` used to create its `dstAccelerationStructure` member **must** be bound to memory that was not allocated with multiple instances
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03776)VUID-vkBuildAccelerationStructuresKHR-pInfos-03776  
  For each element of `pInfos`, if its `mode` member is `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR` the `buffer` used to create its `srcAccelerationStructure` member **must** be bound to memory that was not allocated with multiple instances
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03777)VUID-vkBuildAccelerationStructuresKHR-pInfos-03777  
  For each element of `pInfos`, the `buffer` used to create each acceleration structure referenced by the `geometry.instances.data` member of any element of `pGeometries` or `ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR` **must** be bound to memory that was not allocated with multiple instances
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03778)VUID-vkBuildAccelerationStructuresKHR-pInfos-03778  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR`, `geometry.instances.data.hostAddress` **must** be a valid host address
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-03779)VUID-vkBuildAccelerationStructuresKHR-pInfos-03779  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR`, each [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInstanceKHR.html)::`accelerationStructureReference` value in `geometry.instances.data.hostAddress` **must** be a valid [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html) object
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-04930)VUID-vkBuildAccelerationStructuresKHR-pInfos-04930  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR` with `VK_BUILD_ACCELERATION_STRUCTURE_MOTION_BIT_NV` set, each `accelerationStructureReference` in any structure in [VkAccelerationStructureMotionInstanceNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMotionInstanceNV.html) value in `geometry.instances.data.hostAddress` **must** be a valid [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html) object
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-10892)VUID-vkBuildAccelerationStructuresKHR-pInfos-10892  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries` with a `geometryType` of `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if there is an instance of [VkAccelerationStructureTrianglesOpacityMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureTrianglesOpacityMicromapEXT.html) in the `geometry.triangles.pNext` chain, and its `indexType` is `VK_INDEX_TYPE_NONE_KHR`, then its `indexBuffer.hostAddress` **must** be 0
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-10893)VUID-vkBuildAccelerationStructuresKHR-pInfos-10893  
  For any element of `pInfos`\[i].`pGeometries` or `pInfos`\[i].`ppGeometries`, `geometryType` **must** not be `VK_GEOMETRY_TYPE_DENSE_GEOMETRY_FORMAT_TRIANGLES_AMDX`

Valid Usage (Implicit)

- [](#VUID-vkBuildAccelerationStructuresKHR-device-parameter)VUID-vkBuildAccelerationStructuresKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkBuildAccelerationStructuresKHR-deferredOperation-parameter)VUID-vkBuildAccelerationStructuresKHR-deferredOperation-parameter  
  If `deferredOperation` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `deferredOperation` **must** be a valid [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html) handle
- [](#VUID-vkBuildAccelerationStructuresKHR-pInfos-parameter)VUID-vkBuildAccelerationStructuresKHR-pInfos-parameter  
  `pInfos` **must** be a valid pointer to an array of `infoCount` valid [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html) structures
- [](#VUID-vkBuildAccelerationStructuresKHR-ppBuildRangeInfos-parameter)VUID-vkBuildAccelerationStructuresKHR-ppBuildRangeInfos-parameter  
  `ppBuildRangeInfos` **must** be a valid pointer to an array of `infoCount` [VkAccelerationStructureBuildRangeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildRangeInfoKHR.html) structures
- [](#VUID-vkBuildAccelerationStructuresKHR-infoCount-arraylength)VUID-vkBuildAccelerationStructuresKHR-infoCount-arraylength  
  `infoCount` **must** be greater than `0`
- [](#VUID-vkBuildAccelerationStructuresKHR-deferredOperation-parent)VUID-vkBuildAccelerationStructuresKHR-deferredOperation-parent  
  If `deferredOperation` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_OPERATION_DEFERRED_KHR`
- `VK_OPERATION_NOT_DEFERRED_KHR`
- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html), [VkAccelerationStructureBuildRangeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildRangeInfoKHR.html), [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkBuildAccelerationStructuresKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0