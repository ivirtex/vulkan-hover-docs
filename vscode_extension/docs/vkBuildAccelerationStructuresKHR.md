# vkBuildAccelerationStructuresKHR(3) Manual Page

## Name

vkBuildAccelerationStructuresKHR - Build an acceleration structure on
the host



## <a href="#_c_specification" class="anchor"></a>C Specification

To build acceleration structures on the host, call:

``` c
// Provided by VK_KHR_acceleration_structure
VkResult vkBuildAccelerationStructuresKHR(
    VkDevice                                    device,
    VkDeferredOperationKHR                      deferredOperation,
    uint32_t                                    infoCount,
    const VkAccelerationStructureBuildGeometryInfoKHR* pInfos,
    const VkAccelerationStructureBuildRangeInfoKHR* const* ppBuildRangeInfos);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the `VkDevice` for which the acceleration structures are
  being built.

- `deferredOperation` is an optional
  [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html) to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#deferred-host-operations-requesting"
  target="_blank" rel="noopener">request deferral</a> for this command.

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

This command fulfills the same task as
[vkCmdBuildAccelerationStructuresKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBuildAccelerationStructuresKHR.html)
but is executed by the host.

The `vkBuildAccelerationStructuresKHR` command provides the ability to
initiate multiple acceleration structures builds, however there is no
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
href="vkBuildAccelerationStructuresKHR.html">vkBuildAccelerationStructuresKHR</a>
call as the associated bottom-level or instance acceleration structures
are being built. There also <strong>cannot</strong> be any memory
aliasing between any acceleration structure memories or scratch memories
being used by any of the builds.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a
  href="#VUID-vkBuildAccelerationStructuresKHR-accelerationStructureHostCommands-03581"
  id="VUID-vkBuildAccelerationStructuresKHR-accelerationStructureHostCommands-03581"></a>
  VUID-vkBuildAccelerationStructuresKHR-accelerationStructureHostCommands-03581  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-accelerationStructureHostCommands"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceAccelerationStructureFeaturesKHR</code>::<code>accelerationStructureHostCommands</code></a>
  feature **must** be enabled

<!-- -->

- <a href="#VUID-vkBuildAccelerationStructuresKHR-mode-04628"
  id="VUID-vkBuildAccelerationStructuresKHR-mode-04628"></a>
  VUID-vkBuildAccelerationStructuresKHR-mode-04628  
  The `mode` member of each element of `pInfos` **must** be a valid
  [VkBuildAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildAccelerationStructureModeKHR.html)
  value

- <a
  href="#VUID-vkBuildAccelerationStructuresKHR-srcAccelerationStructure-04629"
  id="VUID-vkBuildAccelerationStructuresKHR-srcAccelerationStructure-04629"></a>
  VUID-vkBuildAccelerationStructuresKHR-srcAccelerationStructure-04629  
  If the `srcAccelerationStructure` member of any element of `pInfos` is
  not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the
  `srcAccelerationStructure` member **must** be a valid
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html) handle

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-04630"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-04630"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-04630  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its
  `srcAccelerationStructure` member **must** not be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03403"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03403"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03403  
  The `srcAccelerationStructure` member of any element of `pInfos`
  **must** not be the same acceleration structure as the
  `dstAccelerationStructure` member of any other element of `pInfos`

- <a
  href="#VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03698"
  id="VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03698"></a>
  VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03698  
  The `dstAccelerationStructure` member of any element of `pInfos`
  **must** not be the same acceleration structure as the
  `dstAccelerationStructure` member of any other element of `pInfos`

- <a
  href="#VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03800"
  id="VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03800"></a>
  VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03800  
  The `dstAccelerationStructure` member of any element of `pInfos`
  **must** be a valid
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html) handle

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03699"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03699"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03699  
  For each element of `pInfos`, if its `type` member is
  `VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR`, its
  `dstAccelerationStructure` member **must** have been created with a
  value of
  [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoKHR.html)::`type`
  equal to either `VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR` or
  `VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR`

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03700"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03700"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03700  
  For each element of `pInfos`, if its `type` member is
  `VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_KHR`, its
  `dstAccelerationStructure` member **must** have been created with a
  value of
  [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoKHR.html)::`type`
  equal to either `VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_KHR` or
  `VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR`

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03663"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03663"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03663  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, [inactive
  primitives](#acceleration-structure-inactive-prims) in its
  `srcAccelerationStructure` member **must** not be made active

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03664"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03664"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03664  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, active primitives
  in its `srcAccelerationStructure` member **must** not be made
  [inactive](#acceleration-structure-inactive-prims)

- <a href="#VUID-vkBuildAccelerationStructuresKHR-None-03407"
  id="VUID-vkBuildAccelerationStructuresKHR-None-03407"></a>
  VUID-vkBuildAccelerationStructuresKHR-None-03407  
  The `dstAccelerationStructure` member of any element of `pInfos`
  **must** not be referenced by the `geometry.instances.data` member of
  any element of `pGeometries` or `ppGeometries` with a `geometryType`
  of `VK_GEOMETRY_TYPE_INSTANCES_KHR` in any other element of `pInfos`

- <a
  href="#VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03701"
  id="VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03701"></a>
  VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03701  
  The range of memory backing the `dstAccelerationStructure` member of
  any element of `pInfos` that is accessed by this command **must** not
  overlap the memory backing the `srcAccelerationStructure` member of
  any other element of `pInfos` with a `mode` equal to
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, which is accessed
  by this command

- <a
  href="#VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03702"
  id="VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03702"></a>
  VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03702  
  The range of memory backing the `dstAccelerationStructure` member of
  any element of `pInfos` that is accessed by this command **must** not
  overlap the memory backing the `dstAccelerationStructure` member of
  any other element of `pInfos`, which is accessed by this command

- <a
  href="#VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03703"
  id="VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03703"></a>
  VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03703  
  The range of memory backing the `dstAccelerationStructure` member of
  any element of `pInfos` that is accessed by this command **must** not
  overlap the memory backing the `scratchData` member of any element of
  `pInfos` (including the same element), which is accessed by this
  command

- <a href="#VUID-vkBuildAccelerationStructuresKHR-scratchData-03704"
  id="VUID-vkBuildAccelerationStructuresKHR-scratchData-03704"></a>
  VUID-vkBuildAccelerationStructuresKHR-scratchData-03704  
  The range of memory backing the `scratchData` member of any element of
  `pInfos` that is accessed by this command **must** not overlap the
  memory backing the `scratchData` member of any other element of
  `pInfos`, which is accessed by this command

- <a href="#VUID-vkBuildAccelerationStructuresKHR-scratchData-03705"
  id="VUID-vkBuildAccelerationStructuresKHR-scratchData-03705"></a>
  VUID-vkBuildAccelerationStructuresKHR-scratchData-03705  
  The range of memory backing the `scratchData` member of any element of
  `pInfos` that is accessed by this command **must** not overlap the
  memory backing the `srcAccelerationStructure` member of any element of
  `pInfos` with a `mode` equal to
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR` (including the same
  element), which is accessed by this command

- <a
  href="#VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03706"
  id="VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03706"></a>
  VUID-vkBuildAccelerationStructuresKHR-dstAccelerationStructure-03706  
  The range of memory backing the `dstAccelerationStructure` member of
  any element of `pInfos` that is accessed by this command **must** not
  overlap the memory backing any acceleration structure referenced by
  the `geometry.instances.data` member of any element of `pGeometries`
  or `ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_INSTANCES_KHR` in any other element of `pInfos`,
  which is accessed by this command

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03667"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03667"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03667  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its
  `srcAccelerationStructure` member **must** have previously been
  constructed with
  `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_UPDATE_BIT_KHR` set in
  [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)::`flags`
  in the build

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03668"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03668"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03668  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its
  `srcAccelerationStructure` and `dstAccelerationStructure` members
  **must** either be the same
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html), or not
  have any [memory aliasing](#resources-memory-aliasing)

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03758"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03758"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03758  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `geometryCount`
  member **must** have the same value which was specified when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03759"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03759"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03759  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `flags` member
  **must** have the same value which was specified when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03760"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03760"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03760  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, its `type` member
  **must** have the same value which was specified when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03761"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03761"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03761  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, its `geometryType` member
  **must** have the same value which was specified when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03762"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03762"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03762  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, its `flags` member **must**
  have the same value which was specified when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03763"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03763"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03763  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, if `geometryType` is
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, its
  `geometry.triangles.vertexFormat` member **must** have the same value
  which was specified when `srcAccelerationStructure` was last built

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03764"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03764"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03764  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, if `geometryType` is
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, its `geometry.triangles.maxVertex`
  member **must** have the same value which was specified when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03765"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03765"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03765  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, if `geometryType` is
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, its `geometry.triangles.indexType`
  member **must** have the same value which was specified when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03766"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03766"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03766  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, if `geometryType` is
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if its
  `geometry.triangles.transformData` address was `NULL` when
  `srcAccelerationStructure` was last built, then it **must** be `NULL`

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03767"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03767"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03767  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, if `geometryType` is
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if its
  `geometry.triangles.transformData` address was not `NULL` when
  `srcAccelerationStructure` was last built, then it **must** not be
  `NULL`

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03768"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03768"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03768  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, if `geometryType` is
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, and `geometry.triangles.indexType`
  is not `VK_INDEX_TYPE_NONE_KHR`, then the value of each index
  referenced **must** be the same as the corresponding index value when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkBuildAccelerationStructuresKHR-primitiveCount-03769"
  id="VUID-vkBuildAccelerationStructuresKHR-primitiveCount-03769"></a>
  VUID-vkBuildAccelerationStructuresKHR-primitiveCount-03769  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, then for each
  `VkAccelerationStructureGeometryKHR` structure referred to by its
  `pGeometries` or `ppGeometries` members, the `primitiveCount` member
  of its corresponding `VkAccelerationStructureBuildRangeInfoKHR`
  structure **must** have the same value which was specified when
  `srcAccelerationStructure` was last built

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03801"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03801"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03801  
  For each element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_INSTANCES_KHR`, the corresponding
  `ppBuildRangeInfos`\[i\]\[j\].`primitiveCount` **must** be less than
  or equal to
  [VkPhysicalDeviceAccelerationStructurePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceAccelerationStructurePropertiesKHR.html)::`maxInstanceCount`

<!-- -->

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03675"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03675"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03675  
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

- <a href="#VUID-vkBuildAccelerationStructuresKHR-ppBuildRangeInfos-03676"
  id="VUID-vkBuildAccelerationStructuresKHR-ppBuildRangeInfos-03676"></a>
  VUID-vkBuildAccelerationStructuresKHR-ppBuildRangeInfos-03676  
  Each element of `ppBuildRangeInfos`\[i\] **must** be a valid pointer
  to an array of `pInfos`\[i\].`geometryCount`
  `VkAccelerationStructureBuildRangeInfoKHR` structures

<!-- -->

- <a href="#VUID-vkBuildAccelerationStructuresKHR-deferredOperation-03678"
  id="VUID-vkBuildAccelerationStructuresKHR-deferredOperation-03678"></a>
  VUID-vkBuildAccelerationStructuresKHR-deferredOperation-03678  
  Any previous deferred operation that was associated with
  `deferredOperation` **must** be complete

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03722"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03722"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03722  
  For each element of `pInfos`, the `buffer` used to create its
  `dstAccelerationStructure` member **must** be bound to host-visible
  device memory

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03723"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03723"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03723  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR` the `buffer` used to
  create its `srcAccelerationStructure` member **must** be bound to
  host-visible device memory

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03724"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03724"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03724  
  For each element of `pInfos`, the `buffer` used to create each
  acceleration structure referenced by the `geometry.instances.data`
  member of any element of `pGeometries` or `ppGeometries` with a
  `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR` **must** be bound
  to host-visible device memory

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03725"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03725"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03725  
  If `pInfos`\[i\].`mode` is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_BUILD_KHR`, all addresses
  between `pInfos`\[i\].`scratchData.hostAddress` and
  `pInfos`\[i\].`scratchData.hostAddress` + N - 1 **must** be valid host
  memory, where N is given by the `buildScratchSize` member of the
  [VkAccelerationStructureBuildSizesInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildSizesInfoKHR.html)
  structure returned from a call to
  [vkGetAccelerationStructureBuildSizesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureBuildSizesKHR.html)
  with an identical
  [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)
  structure and primitive count

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03726"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03726"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03726  
  If `pInfos`\[i\].`mode` is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR`, all addresses
  between `pInfos`\[i\].`scratchData.hostAddress` and
  `pInfos`\[i\].`scratchData.hostAddress` + N - 1 **must** be valid host
  memory, where N is given by the `updateScratchSize` member of the
  [VkAccelerationStructureBuildSizesInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildSizesInfoKHR.html)
  structure returned from a call to
  [vkGetAccelerationStructureBuildSizesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureBuildSizesKHR.html)
  with an identical
  [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)
  structure and primitive count

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03771"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03771"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03771  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`,
  `geometry.triangles.vertexData.hostAddress` **must** be a valid host
  address

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03772"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03772"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03772  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if `geometry.triangles.indexType` is
  not `VK_INDEX_TYPE_NONE_KHR`,
  `geometry.triangles.indexData.hostAddress` **must** be a valid host
  address

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03773"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03773"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03773  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if
  `geometry.triangles.transformData.hostAddress` is not `0`, it **must**
  be a valid host address

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03774"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03774"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03774  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_AABBS_KHR`, `geometry.aabbs.data.hostAddress`
  **must** be a valid host address

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03775"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03775"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03775  
  For each element of `pInfos`, the `buffer` used to create its
  `dstAccelerationStructure` member **must** be bound to memory that was
  not allocated with multiple instances

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03776"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03776"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03776  
  For each element of `pInfos`, if its `mode` member is
  `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR` the `buffer` used to
  create its `srcAccelerationStructure` member **must** be bound to
  memory that was not allocated with multiple instances

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03777"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03777"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03777  
  For each element of `pInfos`, the `buffer` used to create each
  acceleration structure referenced by the `geometry.instances.data`
  member of any element of `pGeometries` or `ppGeometries` with a
  `geometryType` of `VK_GEOMETRY_TYPE_INSTANCES_KHR` **must** be bound
  to memory that was not allocated with multiple instances

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03778"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03778"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03778  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_INSTANCES_KHR`,
  `geometry.instances.data.hostAddress` **must** be a valid host address

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-03779"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-03779"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-03779  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_INSTANCES_KHR`, each
  [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInstanceKHR.html)::`accelerationStructureReference`
  value in `geometry.instances.data.hostAddress` **must** be a valid
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html) object

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-04930"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-04930"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-04930  
  For any element of `pInfos`\[i\].`pGeometries` or
  `pInfos`\[i\].`ppGeometries` with a `geometryType` of
  `VK_GEOMETRY_TYPE_INSTANCES_KHR` with
  `VK_BUILD_ACCELERATION_STRUCTURE_MOTION_BIT_NV` set, each
  `accelerationStructureReference` in any structure in
  [VkAccelerationStructureMotionInstanceNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInstanceNV.html)
  value in `geometry.instances.data.hostAddress` **must** be a valid
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html) object

Valid Usage (Implicit)

- <a href="#VUID-vkBuildAccelerationStructuresKHR-device-parameter"
  id="VUID-vkBuildAccelerationStructuresKHR-device-parameter"></a>
  VUID-vkBuildAccelerationStructuresKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkBuildAccelerationStructuresKHR-deferredOperation-parameter"
  id="VUID-vkBuildAccelerationStructuresKHR-deferredOperation-parameter"></a>
  VUID-vkBuildAccelerationStructuresKHR-deferredOperation-parameter  
  If `deferredOperation` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `deferredOperation` **must** be a valid
  [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html) handle

- <a href="#VUID-vkBuildAccelerationStructuresKHR-pInfos-parameter"
  id="VUID-vkBuildAccelerationStructuresKHR-pInfos-parameter"></a>
  VUID-vkBuildAccelerationStructuresKHR-pInfos-parameter  
  `pInfos` **must** be a valid pointer to an array of `infoCount` valid
  [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)
  structures

- <a
  href="#VUID-vkBuildAccelerationStructuresKHR-ppBuildRangeInfos-parameter"
  id="VUID-vkBuildAccelerationStructuresKHR-ppBuildRangeInfos-parameter"></a>
  VUID-vkBuildAccelerationStructuresKHR-ppBuildRangeInfos-parameter  
  `ppBuildRangeInfos` **must** be a valid pointer to an array of
  `infoCount`
  [VkAccelerationStructureBuildRangeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildRangeInfoKHR.html)
  structures

- <a href="#VUID-vkBuildAccelerationStructuresKHR-infoCount-arraylength"
  id="VUID-vkBuildAccelerationStructuresKHR-infoCount-arraylength"></a>
  VUID-vkBuildAccelerationStructuresKHR-infoCount-arraylength  
  `infoCount` **must** be greater than `0`

- <a
  href="#VUID-vkBuildAccelerationStructuresKHR-deferredOperation-parent"
  id="VUID-vkBuildAccelerationStructuresKHR-deferredOperation-parent"></a>
  VUID-vkBuildAccelerationStructuresKHR-deferredOperation-parent  
  If `deferredOperation` is a valid handle, it **must** have been
  created, allocated, or retrieved from `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_OPERATION_DEFERRED_KHR`

- `VK_OPERATION_NOT_DEFERRED_KHR`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html),
[VkAccelerationStructureBuildRangeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildRangeInfoKHR.html),
[VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkBuildAccelerationStructuresKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
