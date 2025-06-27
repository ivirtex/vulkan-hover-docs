# vkCmdBuildAccelerationStructuresKHR(3) Manual Page

## Name

vkCmdBuildAccelerationStructuresKHR - Build an acceleration structure



## <a href="#_c_specification" class="anchor"></a>C Specification

To build acceleration structures call:

``` c
// Provided by VK_KHR_acceleration_structure
void vkCmdBuildAccelerationStructuresKHR(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    infoCount,
    const VkAccelerationStructureBuildGeometryInfoKHR* pInfos,
    const VkAccelerationStructureBuildRangeInfoKHR* const* ppBuildRangeInfos);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `infoCount` is the number of acceleration structures to build. It
  specifies the number of the `pInfos` structures and
  `ppBuildRangeInfos` pointers that **must** be provided.

- `pInfos` is a pointer to an array of `infoCount`
  [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)
  structures defining the geometry used to build each acceleration
  structure.

- `ppBuildRangeInfos` is a pointer to an array of `infoCount` pointers
  to arrays of
  [VkAccelerationStructureBuildRangeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildRangeInfoKHR.html)
  structures. Each `ppBuildRangeInfos`\[i\] is a pointer to an array of
  `pInfos`\[i\].`geometryCount`
  [VkAccelerationStructureBuildRangeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildRangeInfoKHR.html)
  structures defining dynamic offsets to the addresses where geometry
  data is stored, as defined by `pInfos`\[i\].

## <a href="#_description" class="anchor"></a>Description

The `vkCmdBuildAccelerationStructuresKHR` command provides the ability
to initiate multiple acceleration structures builds, however there is no
ordering or synchronization implied between any of the individual
acceleration structure builds.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>This means that an application <strong>cannot</strong> build a
top-level acceleration structure in the same <a
href="vkCmdBuildAccelerationStructuresKHR.html">vkCmdBuildAccelerationStructuresKHR</a>
call as the associated bottom-level or instance acceleration structures
are being built. There also <strong>cannot</strong> be any memory
aliasing between any acceleration structure memories or scratch memories
being used by any of the builds.</p></td>
</tr>
</tbody>
</table>

Accesses to the acceleration structure scratch buffers as identified by
the
[VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)::`scratchData`
buffer device addresses **must** be <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies"
target="_blank" rel="noopener">synchronized</a> with the
`VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
target="_blank" rel="noopener">pipeline stage</a> and an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-types"
target="_blank" rel="noopener">access type</a> of
(`VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR` \|
`VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_KHR`). Accesses to each
[VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)::`srcAccelerationStructure`
and
[VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)::`dstAccelerationStructure`
**must** be <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies"
target="_blank" rel="noopener">synchronized</a> with the
`VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
target="_blank" rel="noopener">pipeline stage</a> and an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-types"
target="_blank" rel="noopener">access type</a> of
`VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR` or
`VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_KHR`, as appropriate.

Accesses to other input buffers as identified by any used values of
[VkAccelerationStructureGeometryMotionTrianglesDataNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryMotionTrianglesDataNV.html)::`vertexData`,
[VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)::`vertexData`,
[VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)::`indexData`,
[VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)::`transformData`,
[VkAccelerationStructureGeometryAabbsDataKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryAabbsDataKHR.html)::`data`,
and
[VkAccelerationStructureGeometryInstancesDataKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryInstancesDataKHR.html)::`data`
**must** be <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies"
target="_blank" rel="noopener">synchronized</a> with the
`VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
target="_blank" rel="noopener">pipeline stage</a> and an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-types"
target="_blank" rel="noopener">access type</a> of
`VK_ACCESS_SHADER_READ_BIT`.

Valid Usage

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresKHR-accelerationStructure-08923"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-accelerationStructure-08923"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-accelerationStructure-08923  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-accelerationStructure"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceAccelerationStructureFeaturesKHR</code>::<code>accelerationStructure</code></a>
  feature **must** be enabled

<!-- -->

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-mode-04628"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-mode-04628"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-mode-04628  
  The `mode` member of each element of `pInfos` **must** be a valid
  [VkBuildAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildAccelerationStructureModeKHR.html)
  value

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresKHR-srcAccelerationStructure-04629"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-srcAccelerationStructure-04629"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-srcAccelerationStructure-04629  
  If the `srcAccelerationStructure` member of any element of `pInfos` is
  not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the
  `srcAccelerationStructure` member **must** be a valid
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html) handle

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-04630"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-04630"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-04630  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its
  `srcAccelerationStructure` member **must** not be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03403"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03403"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03403  
  The `srcAccelerationStructure` member of any element of `pInfos`
  **must** not be the same acceleration structure as the
  `dstAccelerationStructure` member of any other element of `pInfos`

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03698"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03698"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03698  
  The `dstAccelerationStructure` member of any element of `pInfos`
  **must** not be the same acceleration structure as the
  `dstAccelerationStructure` member of any other element of `pInfos`

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03800"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03800"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03800  
  The `dstAccelerationStructure` member of any element of `pInfos`
  **must** be a valid
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html) handle

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03699"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03699"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03699  
  For each element of `pInfos`, if its `type` member is
  `VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR`, its
  `dstAccelerationStructure` member **must** have been created with a
  value of
  [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoKHR.html)::`type`
  equal to either `VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR` or
  `VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR`

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03700"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03700"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03700  
  For each element of `pInfos`, if its `type` member is
  `VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_KHR`, its
  `dstAccelerationStructure` member **must** have been created with a
  value of
  [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoKHR.html)::`type`
  equal to either `VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_KHR` or
  `VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR`

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03663"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03663"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03663  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, [inactive
  primitives](#acceleration-structure-inactive-prims) in its
  `srcAccelerationStructure` member **must** not be made active

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03664"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03664"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03664  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, active primitives
  in its `srcAccelerationStructure` member **must** not be made
  [inactive](#acceleration-structure-inactive-prims)

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-None-03407"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-None-03407"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-None-03407  
  The `dstAccelerationStructure` member of any element of `pInfos`
  **must** not be referenced by the `geometry.instances.data` member of
  any element of `pGeometries` or `ppGeometries` with a `geometryType`
  of `VK_GEOMETRY_TYPE_INSTANCES_KHR` in any other element of `pInfos`

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03701"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03701"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03701  
  The range of memory backing the `dstAccelerationStructure` member of
  any element of `pInfos` that is accessed by this command **must** not
  overlap the memory backing the `srcAccelerationStructure` member of
  any other element of `pInfos` with a `mode` equal to
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, which is accessed
  by this command

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03702"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03702"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03702  
  The range of memory backing the `dstAccelerationStructure` member of
  any element of `pInfos` that is accessed by this command **must** not
  overlap the memory backing the `dstAccelerationStructure` member of
  any other element of `pInfos`, which is accessed by this command

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03703"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03703"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03703  
  The range of memory backing the `dstAccelerationStructure` member of
  any element of `pInfos` that is accessed by this command **must** not
  overlap the memory backing the `scratchData` member of any element of
  `pInfos` (including the same element), which is accessed by this
  command

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-scratchData-03704"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-scratchData-03704"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-scratchData-03704  
  The range of memory backing the `scratchData` member of any element of
  `pInfos` that is accessed by this command **must** not overlap the
  memory backing the `scratchData` member of any other element of
  `pInfos`, which is accessed by this command

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-scratchData-03705"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-scratchData-03705"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-scratchData-03705  
  The range of memory backing the `scratchData` member of any element of
  `pInfos` that is accessed by this command **must** not overlap the
  memory backing the `srcAccelerationStructure` member of any element of
  `pInfos` with a `mode` equal to
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR` (including the same
  element), which is accessed by this command

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03706"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03706"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-dstAccelerationStructure-03706  
  The range of memory backing the `dstAccelerationStructure` member of
  any element of `pInfos` that is accessed by this command **must** not
  overlap the memory backing any acceleration structure referenced by
  the `geometry.instances.data` member of any element of `pGeometries`
  or `ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_INSTANCES_KHR` in any other element of `pInfos`,
  which is accessed by this command

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03667"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03667"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03667  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its
  `srcAccelerationStructure` member **must** have previously been
  constructed with
  `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_UPDATE_BIT_KHR` set in
  [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)::`flags`
  in the build

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03668"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03668"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03668  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its
  `srcAccelerationStructure` and `dstAccelerationStructure` members
  **must** either be the same
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html), or not
  have any [memory aliasing](#resources-memory-aliasing)

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03758"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03758"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03758  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `geometryCount`
  member **must** have the same value which was specified when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03759"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03759"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03759  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `flags` member
  **must** have the same value which was specified when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03760"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03760"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03760  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `type` member
  **must** have the same value which was specified when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03761"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03761"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03761  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, its `geometryType` member
  **must** have the same value which was specified when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03762"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03762"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03762  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, its `flags` member **must**
  have the same value which was specified when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03763"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03763"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03763  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, if `geometryType` is
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, its
  `geometry.triangles.vertexFormat` member **must** have the same value
  which was specified when `srcAccelerationStructure` was last built

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03764"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03764"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03764  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, if `geometryType` is
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, its `geometry.triangles.maxVertex`
  member **must** have the same value which was specified when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03765"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03765"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03765  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, if `geometryType` is
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, its `geometry.triangles.indexType`
  member **must** have the same value which was specified when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03766"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03766"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03766  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, if `geometryType` is
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if its
  `geometry.triangles.transformData` address was `NULL` when
  `srcAccelerationStructure` was last built, then it **must** be `NULL`

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03767"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03767"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03767  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, if `geometryType` is
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if its
  `geometry.triangles.transformData` address was not `NULL` when
  `srcAccelerationStructure` was last built, then it **must** not be
  `NULL`

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03768"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03768"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03768  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, if `geometryType` is
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, and `geometry.triangles.indexType`
  is not `VK_INDEX_TYPE_NONE_KHR`, then the value of each index
  referenced **must** be the same as the corresponding index value when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-primitiveCount-03769"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-primitiveCount-03769"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-primitiveCount-03769  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, the `primitiveCount` member
  of its corresponding `VkAccelerationStructureBuildRangeInfoKHR`
  structure **must** have the same value which was specified when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03801"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03801"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03801  
  For each element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_INSTANCES_KHR`, the corresponding
  `ppBuildRangeInfos`\[i\]\[j\].`primitiveCount` **must** be less than
  or equal to
  [VkPhysicalDeviceAccelerationStructurePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceAccelerationStructurePropertiesKHR.html)::`maxInstanceCount`

<!-- -->

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03707"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03707"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03707  
  For each element of `pInfos`, the `buffer` used to create its
  `dstAccelerationStructure` member **must** be bound to device memory

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03708"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03708"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03708  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR` the `buffer` used to
  create its `srcAccelerationStructure` member **must** be bound to
  device memory

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03709"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03709"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03709  
  For each element of `pInfos`, the `buffer` used to create each
  acceleration structure referenced by the `geometry.instances.data`
  member of any element of `pGeometries` or `ppGeometries` with a
  `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR` **must** be bound
  to device memory

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03671"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03671"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03671  
  If `pInfos`\[i\].`mode` is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_BUILD_KHR`, all addresses
  between `pInfos`\[i\].`scratchData.deviceAddress` and
  `pInfos`\[i\].`scratchData.deviceAddress` + N - 1 **must** be in the
  buffer device address range of the same buffer, where N is given by
  the `buildScratchSize` member of the
  [VkAccelerationStructureBuildSizesInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildSizesInfoKHR.html)
  structure returned from a call to
  [vkGetAccelerationStructureBuildSizesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureBuildSizesKHR.html)
  with an identical
  [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)
  structure and primitive count

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03672"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03672"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03672  
  If `pInfos`\[i\].`mode` is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, all addresses
  between `pInfos`\[i\].`scratchData.deviceAddress` and
  `pInfos`\[i\].`scratchData.deviceAddress` + N - 1 **must** be in the
  buffer device address range of the same buffer, where N is given by
  the `updateScratchSize` member of the
  [VkAccelerationStructureBuildSizesInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildSizesInfoKHR.html)
  structure returned from a call to
  [vkGetAccelerationStructureBuildSizesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureBuildSizesKHR.html)
  with an identical
  [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)
  structure and primitive count

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-geometry-03673"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-geometry-03673"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-geometry-03673  
  The buffers from which the buffer device addresses for all of the
  `geometry.triangles.vertexData`, `geometry.triangles.indexData`,
  `geometry.triangles.transformData`, `geometry.aabbs.data`, and
  `geometry.instances.data` members of all `pInfos`\[i\].`pGeometries`
  and `pInfos`\[i\].`ppGeometries` are queried **must** have been
  created with the
  `VK_BUFFER_USAGE_ACCELERATION_STRUCTURE_BUILD_INPUT_READ_ONLY_BIT_KHR`
  usage flag

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03674"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03674"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03674  
  The buffer from which the buffer device address
  `pInfos`\[i\].`scratchData.deviceAddress` is queried **must** have
  been created with `VK_BUFFER_USAGE_STORAGE_BUFFER_BIT` usage flag

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03802"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03802"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03802  
  For each element of `pInfos`, its `scratchData.deviceAddress` member
  **must** be a valid device address obtained from
  [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html)

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03803"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03803"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03803  
  For each element of `pInfos`, if `scratchData.deviceAddress` is the
  address of a non-sparse buffer then it **must** be bound completely
  and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)
  object

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03710"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03710"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03710  
  For each element of `pInfos`, its `scratchData.deviceAddress` member
  **must** be a multiple of
  [VkPhysicalDeviceAccelerationStructurePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceAccelerationStructurePropertiesKHR.html)::`minAccelerationStructureScratchOffsetAlignment`

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03804"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03804"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03804  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`,
  `geometry.triangles.vertexData.deviceAddress` **must** be a valid
  device address obtained from
  [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html)

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03805"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03805"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03805  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if
  `geometry.triangles.vertexData.deviceAddress` is the address of a
  non-sparse buffer then it **must** be bound completely and
  contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03711"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03711"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03711  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`,
  `geometry.triangles.vertexData.deviceAddress` **must** be aligned to
  the size in bytes of the smallest component of the format in
  `vertexFormat`

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03806"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03806"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03806  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if `geometry.triangles.indexType` is
  not `VK_INDEX_TYPE_NONE_KHR`,
  `geometry.triangles.indexData.deviceAddress` **must** be a valid
  device address obtained from
  [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html)

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03807"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03807"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03807  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if `geometry.triangles.indexType` is
  not `VK_INDEX_TYPE_NONE_KHR`, if
  `geometry.triangles.indexData.deviceAddress` is the address of a
  non-sparse buffer then it **must** be bound completely and
  contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03712"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03712"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03712  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, and with
  `geometry.triangles.indexType` not equal to `VK_INDEX_TYPE_NONE_KHR`,
  `geometry.triangles.indexData.deviceAddress` **must** be aligned to
  the size in bytes of the type in `indexType`

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03808"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03808"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03808  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if
  `geometry.triangles.transformData.deviceAddress` is not `0`, it
  **must** be a valid device address obtained from
  [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html)

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03809"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03809"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03809  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if
  `geometry.triangles.transformData.deviceAddress` is the address of a
  non-sparse buffer then it **must** be bound completely and
  contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03810"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03810"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03810  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if
  `geometry.triangles.transformData.deviceAddress` is not `0`, it
  **must** be aligned to `16` bytes

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03811"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03811"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03811  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_AABBS_KHR`, `geometry.aabbs.data.deviceAddress`
  **must** be a valid device address obtained from
  [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html)

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03812"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03812"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03812  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_AABBS_KHR`, if `geometry.aabbs.data.deviceAddress`
  is the address of a non-sparse buffer then it **must** be bound
  completely and contiguously to a single
  [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03714"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03714"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03714  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_AABBS_KHR`, `geometry.aabbs.data.deviceAddress`
  **must** be aligned to `8` bytes

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03715"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03715"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03715  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_INSTANCES_KHR`, if `geometry.arrayOfPointers` is
  `VK_FALSE`, `geometry.instances.data.deviceAddress` **must** be
  aligned to `16` bytes

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03716"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03716"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03716  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_INSTANCES_KHR`, if `geometry.arrayOfPointers` is
  `VK_TRUE`, `geometry.instances.data.deviceAddress` **must** be aligned
  to `8` bytes

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03717"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03717"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03717  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_INSTANCES_KHR`, if `geometry.arrayOfPointers` is
  `VK_TRUE`, each element of `geometry.instances.data.deviceAddress` in
  device memory **must** be aligned to `16` bytes

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03813"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03813"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03813  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_INSTANCES_KHR`,
  `geometry.instances.data.deviceAddress` **must** be a valid device
  address obtained from
  [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html)

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03814"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03814"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03814  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_INSTANCES_KHR`, if
  `geometry.instances.data.deviceAddress` is the address of a non-sparse
  buffer then it **must** be bound completely and contiguously to a
  single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-06707"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-06707"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-06707  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_INSTANCES_KHR`, each
  [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInstanceKHR.html)::`accelerationStructureReference`
  value in `geometry.instances.data.deviceAddress` **must** be a valid
  device address containing a value obtained from
  [vkGetAccelerationStructureDeviceAddressKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureDeviceAddressKHR.html)
  or `0`

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-commandBuffer-09547"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-commandBuffer-09547"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-commandBuffer-09547  
  `commandBuffer` **must** not be a protected command buffer

<!-- -->

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03675"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03675"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-03675  
  For each `pInfos`\[i\], `dstAccelerationStructure` **must** have been
  created with a value of
  [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoKHR.html)::`size`
  greater than or equal to the memory size required by the build
  operation, as returned by
  [vkGetAccelerationStructureBuildSizesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureBuildSizesKHR.html)
  with `pBuildInfo` = `pInfos`\[i\] and with each element of the
  `pMaxPrimitiveCounts` array greater than or equal to the equivalent
  `ppBuildRangeInfos`\[i\]\[j\].`primitiveCount` values for `j` in
  \[0,`pInfos`\[i\].`geometryCount`)

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresKHR-ppBuildRangeInfos-03676"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-ppBuildRangeInfos-03676"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-ppBuildRangeInfos-03676  
  Each element of `ppBuildRangeInfos`\[i\] **must** be a valid pointer
  to an array of `pInfos`\[i\].`geometryCount`
  `VkAccelerationStructureBuildRangeInfoKHR` structures

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresKHR-commandBuffer-parameter"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-commandBuffer-parameter"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-parameter"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-parameter"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-pInfos-parameter  
  `pInfos` **must** be a valid pointer to an array of `infoCount` valid
  [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)
  structures

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresKHR-ppBuildRangeInfos-parameter"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-ppBuildRangeInfos-parameter"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-ppBuildRangeInfos-parameter  
  `ppBuildRangeInfos` **must** be a valid pointer to an array of
  `infoCount`
  [VkAccelerationStructureBuildRangeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildRangeInfoKHR.html)
  structures

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresKHR-commandBuffer-recording"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-commandBuffer-recording"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresKHR-commandBuffer-cmdpool"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-commandBuffer-cmdpool"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support compute operations

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-renderpass"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-renderpass"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdBuildAccelerationStructuresKHR-videocoding"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-videocoding"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-videocoding  
  This command **must** only be called outside of a video coding scope

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresKHR-infoCount-arraylength"
  id="VUID-vkCmdBuildAccelerationStructuresKHR-infoCount-arraylength"></a>
  VUID-vkCmdBuildAccelerationStructuresKHR-infoCount-arraylength  
  `infoCount` **must** be greater than `0`

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized

- Host access to the `VkCommandPool` that `commandBuffer` was allocated
  from **must** be externally synchronized

Command Properties

<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-top"><a
href="#VkCommandBufferLevel">Command Buffer Levels</a></th>
<th class="tableblock halign-left valign-top"><a
href="#vkCmdBeginRenderPass">Render Pass Scope</a></th>
<th class="tableblock halign-left valign-top"><a
href="#vkCmdBeginVideoCodingKHR">Video Coding Scope</a></th>
<th class="tableblock halign-left valign-top"><a
href="#VkQueueFlagBits">Supported Queue Types</a></th>
<th class="tableblock halign-left valign-top"><a
href="#fundamentals-queueoperation-command-types">Command Type</a></th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p>Primary<br />
Secondary</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Compute</p></td>
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html),
[VkAccelerationStructureBuildRangeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildRangeInfoKHR.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBuildAccelerationStructuresKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
