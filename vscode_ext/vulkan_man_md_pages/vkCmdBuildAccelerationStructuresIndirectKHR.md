# vkCmdBuildAccelerationStructuresIndirectKHR(3) Manual Page

## Name

vkCmdBuildAccelerationStructuresIndirectKHR - Build an acceleration
structure with some parameters provided on the device



## <a href="#_c_specification" class="anchor"></a>C Specification

To build acceleration structures with some parameters sourced on the
device call:

``` c
// Provided by VK_KHR_acceleration_structure
void vkCmdBuildAccelerationStructuresIndirectKHR(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    infoCount,
    const VkAccelerationStructureBuildGeometryInfoKHR* pInfos,
    const VkDeviceAddress*                      pIndirectDeviceAddresses,
    const uint32_t*                             pIndirectStrides,
    const uint32_t* const*                      ppMaxPrimitiveCounts);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `infoCount` is the number of acceleration structures to build.

- `pInfos` is a pointer to an array of `infoCount`
  [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)
  structures defining the geometry used to build each acceleration
  structure.

- `pIndirectDeviceAddresses` is a pointer to an array of `infoCount`
  buffer device addresses which point to `pInfos`\[i\].`geometryCount`
  [VkAccelerationStructureBuildRangeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildRangeInfoKHR.html)
  structures defining dynamic offsets to the addresses where geometry
  data is stored, as defined by `pInfos`\[i\].

- `pIndirectStrides` is a pointer to an array of `infoCount` byte
  strides between elements of `pIndirectDeviceAddresses`.

- `ppMaxPrimitiveCounts` is a pointer to an array of `infoCount`
  pointers to arrays of `pInfos`\[i\].`geometryCount` values indicating
  the maximum number of primitives that will be built by this command
  for each geometry.

## <a href="#_description" class="anchor"></a>Description

Accesses to acceleration structures, scratch buffers, vertex buffers,
index buffers, and instance buffers must be synchronized as with <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#acceleration-structure-scratch"
target="_blank" rel="noopener">vkCmdBuildAccelerationStructuresKHR</a>.

Accesses to any element of `pIndirectDeviceAddresses` **must** be <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies"
target="_blank" rel="noopener">synchronized</a> with the
`VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
target="_blank" rel="noopener">pipeline stage</a> and an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-types"
target="_blank" rel="noopener">access type</a> of
`VK_ACCESS_INDIRECT_COMMAND_READ_BIT`.

Valid Usage

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-accelerationStructureIndirectBuild-03650"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-accelerationStructureIndirectBuild-03650"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-accelerationStructureIndirectBuild-03650  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-accelerationStructureIndirectBuild"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceAccelerationStructureFeaturesKHR</code>::<code>accelerationStructureIndirectBuild</code></a>
  feature **must** be enabled

<!-- -->

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-mode-04628"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-mode-04628"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-mode-04628  
  The `mode` member of each element of `pInfos` **must** be a valid
  [VkBuildAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildAccelerationStructureModeKHR.html)
  value

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-srcAccelerationStructure-04629"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-srcAccelerationStructure-04629"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-srcAccelerationStructure-04629  
  If the `srcAccelerationStructure` member of any element of `pInfos` is
  not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the
  `srcAccelerationStructure` member **must** be a valid
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html) handle

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-04630"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-04630"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-04630  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its
  `srcAccelerationStructure` member **must** not be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03403"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03403"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03403  
  The `srcAccelerationStructure` member of any element of `pInfos`
  **must** not be the same acceleration structure as the
  `dstAccelerationStructure` member of any other element of `pInfos`

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03698"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03698"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03698  
  The `dstAccelerationStructure` member of any element of `pInfos`
  **must** not be the same acceleration structure as the
  `dstAccelerationStructure` member of any other element of `pInfos`

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03800"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03800"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03800  
  The `dstAccelerationStructure` member of any element of `pInfos`
  **must** be a valid
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html) handle

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03699"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03699"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03699  
  For each element of `pInfos`, if its `type` member is
  `VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR`, its
  `dstAccelerationStructure` member **must** have been created with a
  value of
  [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoKHR.html)::`type`
  equal to either `VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR` or
  `VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR`

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03700"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03700"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03700  
  For each element of `pInfos`, if its `type` member is
  `VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_KHR`, its
  `dstAccelerationStructure` member **must** have been created with a
  value of
  [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoKHR.html)::`type`
  equal to either `VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_KHR` or
  `VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR`

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03663"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03663"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03663  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, [inactive
  primitives](#acceleration-structure-inactive-prims) in its
  `srcAccelerationStructure` member **must** not be made active

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03664"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03664"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03664  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, active primitives
  in its `srcAccelerationStructure` member **must** not be made
  [inactive](#acceleration-structure-inactive-prims)

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-None-03407"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-None-03407"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-None-03407  
  The `dstAccelerationStructure` member of any element of `pInfos`
  **must** not be referenced by the `geometry.instances.data` member of
  any element of `pGeometries` or `ppGeometries` with a `geometryType`
  of `VK_GEOMETRY_TYPE_INSTANCES_KHR` in any other element of `pInfos`

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03701"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03701"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03701  
  The range of memory backing the `dstAccelerationStructure` member of
  any element of `pInfos` that is accessed by this command **must** not
  overlap the memory backing the `srcAccelerationStructure` member of
  any other element of `pInfos` with a `mode` equal to
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, which is accessed
  by this command

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03702"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03702"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03702  
  The range of memory backing the `dstAccelerationStructure` member of
  any element of `pInfos` that is accessed by this command **must** not
  overlap the memory backing the `dstAccelerationStructure` member of
  any other element of `pInfos`, which is accessed by this command

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03703"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03703"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03703  
  The range of memory backing the `dstAccelerationStructure` member of
  any element of `pInfos` that is accessed by this command **must** not
  overlap the memory backing the `scratchData` member of any element of
  `pInfos` (including the same element), which is accessed by this
  command

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-scratchData-03704"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-scratchData-03704"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-scratchData-03704  
  The range of memory backing the `scratchData` member of any element of
  `pInfos` that is accessed by this command **must** not overlap the
  memory backing the `scratchData` member of any other element of
  `pInfos`, which is accessed by this command

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-scratchData-03705"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-scratchData-03705"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-scratchData-03705  
  The range of memory backing the `scratchData` member of any element of
  `pInfos` that is accessed by this command **must** not overlap the
  memory backing the `srcAccelerationStructure` member of any element of
  `pInfos` with a `mode` equal to
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR` (including the same
  element), which is accessed by this command

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03706"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03706"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-dstAccelerationStructure-03706  
  The range of memory backing the `dstAccelerationStructure` member of
  any element of `pInfos` that is accessed by this command **must** not
  overlap the memory backing any acceleration structure referenced by
  the `geometry.instances.data` member of any element of `pGeometries`
  or `ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_INSTANCES_KHR` in any other element of `pInfos`,
  which is accessed by this command

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03667"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03667"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03667  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its
  `srcAccelerationStructure` member **must** have previously been
  constructed with
  `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_UPDATE_BIT_KHR` set in
  [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)::`flags`
  in the build

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03668"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03668"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03668  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its
  `srcAccelerationStructure` and `dstAccelerationStructure` members
  **must** either be the same
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html), or not
  have any [memory aliasing](#resources-memory-aliasing)

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03758"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03758"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03758  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `geometryCount`
  member **must** have the same value which was specified when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03759"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03759"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03759  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `flags` member
  **must** have the same value which was specified when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03760"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03760"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03760  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `type` member
  **must** have the same value which was specified when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03761"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03761"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03761  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, its `geometryType` member
  **must** have the same value which was specified when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03762"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03762"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03762  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, its `flags` member **must**
  have the same value which was specified when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03763"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03763"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03763  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, if `geometryType` is
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, its
  `geometry.triangles.vertexFormat` member **must** have the same value
  which was specified when `srcAccelerationStructure` was last built

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03764"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03764"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03764  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, if `geometryType` is
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, its `geometry.triangles.maxVertex`
  member **must** have the same value which was specified when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03765"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03765"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03765  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, if `geometryType` is
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, its `geometry.triangles.indexType`
  member **must** have the same value which was specified when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03766"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03766"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03766  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, if `geometryType` is
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if its
  `geometry.triangles.transformData` address was `NULL` when
  `srcAccelerationStructure` was last built, then it **must** be `NULL`

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03767"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03767"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03767  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, if `geometryType` is
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if its
  `geometry.triangles.transformData` address was not `NULL` when
  `srcAccelerationStructure` was last built, then it **must** not be
  `NULL`

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03768"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03768"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03768  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, if `geometryType` is
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, and `geometry.triangles.indexType`
  is not `VK_INDEX_TYPE_NONE_KHR`, then the value of each index
  referenced **must** be the same as the corresponding index value when
  `srcAccelerationStructure` was last built

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-primitiveCount-03769"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-primitiveCount-03769"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-primitiveCount-03769  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, the `primitiveCount` member
  of its corresponding `VkAccelerationStructureBuildRangeInfoKHR`
  structure **must** have the same value which was specified when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03801"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03801"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03801  
  For each element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_INSTANCES_KHR`, the corresponding
  `ppMaxPrimitiveCounts`\[i\]\[j\] **must** be less than or equal to
  [VkPhysicalDeviceAccelerationStructurePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceAccelerationStructurePropertiesKHR.html)::`maxInstanceCount`

<!-- -->

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03707"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03707"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03707  
  For each element of `pInfos`, the `buffer` used to create its
  `dstAccelerationStructure` member **must** be bound to device memory

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03708"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03708"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03708  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR` the `buffer` used to
  create its `srcAccelerationStructure` member **must** be bound to
  device memory

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03709"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03709"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03709  
  For each element of `pInfos`, the `buffer` used to create each
  acceleration structure referenced by the `geometry.instances.data`
  member of any element of `pGeometries` or `ppGeometries` with a
  `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR` **must** be bound
  to device memory

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03671"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03671"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03671  
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

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03672"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03672"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03672  
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

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-geometry-03673"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-geometry-03673"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-geometry-03673  
  The buffers from which the buffer device addresses for all of the
  `geometry.triangles.vertexData`, `geometry.triangles.indexData`,
  `geometry.triangles.transformData`, `geometry.aabbs.data`, and
  `geometry.instances.data` members of all `pInfos`\[i\].`pGeometries`
  and `pInfos`\[i\].`ppGeometries` are queried **must** have been
  created with the
  `VK_BUFFER_USAGE_ACCELERATION_STRUCTURE_BUILD_INPUT_READ_ONLY_BIT_KHR`
  usage flag

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03674"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03674"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03674  
  The buffer from which the buffer device address
  `pInfos`\[i\].`scratchData.deviceAddress` is queried **must** have
  been created with `VK_BUFFER_USAGE_STORAGE_BUFFER_BIT` usage flag

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03802"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03802"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03802  
  For each element of `pInfos`, its `scratchData.deviceAddress` member
  **must** be a valid device address obtained from
  [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html)

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03803"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03803"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03803  
  For each element of `pInfos`, if `scratchData.deviceAddress` is the
  address of a non-sparse buffer then it **must** be bound completely
  and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)
  object

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03710"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03710"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03710  
  For each element of `pInfos`, its `scratchData.deviceAddress` member
  **must** be a multiple of
  [VkPhysicalDeviceAccelerationStructurePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceAccelerationStructurePropertiesKHR.html)::`minAccelerationStructureScratchOffsetAlignment`

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03804"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03804"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03804  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`,
  `geometry.triangles.vertexData.deviceAddress` **must** be a valid
  device address obtained from
  [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html)

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03805"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03805"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03805  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if
  `geometry.triangles.vertexData.deviceAddress` is the address of a
  non-sparse buffer then it **must** be bound completely and
  contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03711"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03711"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03711  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`,
  `geometry.triangles.vertexData.deviceAddress` **must** be aligned to
  the size in bytes of the smallest component of the format in
  `vertexFormat`

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03806"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03806"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03806  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if `geometry.triangles.indexType` is
  not `VK_INDEX_TYPE_NONE_KHR`,
  `geometry.triangles.indexData.deviceAddress` **must** be a valid
  device address obtained from
  [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html)

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03807"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03807"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03807  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if `geometry.triangles.indexType` is
  not `VK_INDEX_TYPE_NONE_KHR`, if
  `geometry.triangles.indexData.deviceAddress` is the address of a
  non-sparse buffer then it **must** be bound completely and
  contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03712"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03712"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03712  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, and with
  `geometry.triangles.indexType` not equal to `VK_INDEX_TYPE_NONE_KHR`,
  `geometry.triangles.indexData.deviceAddress` **must** be aligned to
  the size in bytes of the type in `indexType`

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03808"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03808"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03808  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if
  `geometry.triangles.transformData.deviceAddress` is not `0`, it
  **must** be a valid device address obtained from
  [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html)

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03809"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03809"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03809  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if
  `geometry.triangles.transformData.deviceAddress` is the address of a
  non-sparse buffer then it **must** be bound completely and
  contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03810"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03810"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03810  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if
  `geometry.triangles.transformData.deviceAddress` is not `0`, it
  **must** be aligned to `16` bytes

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03811"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03811"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03811  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_AABBS_KHR`, `geometry.aabbs.data.deviceAddress`
  **must** be a valid device address obtained from
  [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html)

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03812"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03812"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03812  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_AABBS_KHR`, if `geometry.aabbs.data.deviceAddress`
  is the address of a non-sparse buffer then it **must** be bound
  completely and contiguously to a single
  [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03714"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03714"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03714  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_AABBS_KHR`, `geometry.aabbs.data.deviceAddress`
  **must** be aligned to `8` bytes

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03715"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03715"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03715  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_INSTANCES_KHR`, if `geometry.arrayOfPointers` is
  `VK_FALSE`, `geometry.instances.data.deviceAddress` **must** be
  aligned to `16` bytes

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03716"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03716"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03716  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_INSTANCES_KHR`, if `geometry.arrayOfPointers` is
  `VK_TRUE`, `geometry.instances.data.deviceAddress` **must** be aligned
  to `8` bytes

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03717"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03717"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03717  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_INSTANCES_KHR`, if `geometry.arrayOfPointers` is
  `VK_TRUE`, each element of `geometry.instances.data.deviceAddress` in
  device memory **must** be aligned to `16` bytes

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03813"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03813"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03813  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_INSTANCES_KHR`,
  `geometry.instances.data.deviceAddress` **must** be a valid device
  address obtained from
  [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html)

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03814"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03814"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03814  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_INSTANCES_KHR`, if
  `geometry.instances.data.deviceAddress` is the address of a non-sparse
  buffer then it **must** be bound completely and contiguously to a
  single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-06707"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-06707"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-06707  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_INSTANCES_KHR`, each
  [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInstanceKHR.html)::`accelerationStructureReference`
  value in `geometry.instances.data.deviceAddress` **must** be a valid
  device address containing a value obtained from
  [vkGetAccelerationStructureDeviceAddressKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureDeviceAddressKHR.html)
  or `0`

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-commandBuffer-09547"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-commandBuffer-09547"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-commandBuffer-09547  
  `commandBuffer` **must** not be a protected command buffer

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-03645"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-03645"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-03645  
  For any element of `pIndirectDeviceAddresses`, if the buffer from
  which it was queried is non-sparse then it **must** be bound
  completely and contiguously to a single
  [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-03646"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-03646"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-03646  
  For any element of `pIndirectDeviceAddresses`\[i\], all device
  addresses between `pIndirectDeviceAddresses`\[i\] and
  `pIndirectDeviceAddresses`\[i\] + (`pInfos`\[i\].`geometryCount` ×
  `pIndirectStrides`\[i\]) - 1 **must** be in the buffer device address
  range of the same buffer

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-03647"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-03647"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-03647  
  For any element of `pIndirectDeviceAddresses`, the buffer from which
  it was queried **must** have been created with the
  `VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT` bit set

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-03648"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-03648"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-03648  
  Each element of `pIndirectDeviceAddresses` **must** be a multiple of
  `4`

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectStrides-03787"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectStrides-03787"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectStrides-03787  
  Each element of `pIndirectStrides` **must** be a multiple of `4`

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-03651"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-03651"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-03651  
  Each
  [VkAccelerationStructureBuildRangeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildRangeInfoKHR.html)
  structure referenced by any element of `pIndirectDeviceAddresses`
  **must** be a valid
  [VkAccelerationStructureBuildRangeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildRangeInfoKHR.html)
  structure

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03652"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03652"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-03652  
  `pInfos`\[i\].`dstAccelerationStructure` **must** have been created
  with a value of
  [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoKHR.html)::`size`
  greater than or equal to the memory size required by the build
  operation, as returned by
  [vkGetAccelerationStructureBuildSizesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureBuildSizesKHR.html)
  with `pBuildInfo` = `pInfos`\[i\] and `pMaxPrimitiveCounts` =
  `ppMaxPrimitiveCounts`\[i\]

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-ppMaxPrimitiveCounts-03653"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-ppMaxPrimitiveCounts-03653"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-ppMaxPrimitiveCounts-03653  
  Each `ppMaxPrimitiveCounts`\[i\]\[j\] **must** be greater than or
  equal to the `primitiveCount` value specified by the
  [VkAccelerationStructureBuildRangeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildRangeInfoKHR.html)
  structure located at `pIndirectDeviceAddresses`\[i\] + (`j` ×
  `pIndirectStrides`\[i\])

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-commandBuffer-parameter"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-commandBuffer-parameter"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-parameter"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-parameter"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pInfos-parameter  
  `pInfos` **must** be a valid pointer to an array of `infoCount` valid
  [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)
  structures

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-parameter"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-parameter"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectDeviceAddresses-parameter  
  `pIndirectDeviceAddresses` **must** be a valid pointer to an array of
  `infoCount` [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html) values

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectStrides-parameter"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectStrides-parameter"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-pIndirectStrides-parameter  
  `pIndirectStrides` **must** be a valid pointer to an array of
  `infoCount` `uint32_t` values

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-ppMaxPrimitiveCounts-parameter"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-ppMaxPrimitiveCounts-parameter"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-ppMaxPrimitiveCounts-parameter  
  `ppMaxPrimitiveCounts` **must** be a valid pointer to an array of
  `infoCount` `uint32_t` values

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-commandBuffer-recording"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-commandBuffer-recording"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-commandBuffer-cmdpool"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-commandBuffer-cmdpool"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support compute operations

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-renderpass"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-renderpass"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-videocoding"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-videocoding"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-videocoding  
  This command **must** only be called outside of a video coding scope

- <a
  href="#VUID-vkCmdBuildAccelerationStructuresIndirectKHR-infoCount-arraylength"
  id="VUID-vkCmdBuildAccelerationStructuresIndirectKHR-infoCount-arraylength"></a>
  VUID-vkCmdBuildAccelerationStructuresIndirectKHR-infoCount-arraylength  
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
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBuildAccelerationStructuresIndirectKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
