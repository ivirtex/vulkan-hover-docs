# vkCmdBuildAccelerationStructureNV(3) Manual Page

## Name

vkCmdBuildAccelerationStructureNV - Build an acceleration structure



## <a href="#_c_specification" class="anchor"></a>C Specification

To build an acceleration structure call:

``` c
// Provided by VK_NV_ray_tracing
void vkCmdBuildAccelerationStructureNV(
    VkCommandBuffer                             commandBuffer,
    const VkAccelerationStructureInfoNV*        pInfo,
    VkBuffer                                    instanceData,
    VkDeviceSize                                instanceOffset,
    VkBool32                                    update,
    VkAccelerationStructureNV                   dst,
    VkAccelerationStructureNV                   src,
    VkBuffer                                    scratch,
    VkDeviceSize                                scratchOffset);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `pInfo` contains the shared information for the acceleration
  structure’s structure.

- `instanceData` is the buffer containing an array of
  [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInstanceKHR.html)
  structures defining acceleration structures. This parameter **must**
  be `NULL` for bottom level acceleration structures.

- `instanceOffset` is the offset in bytes (relative to the start of
  `instanceData`) at which the instance data is located.

- `update` specifies whether to update the `dst` acceleration structure
  with the data in `src`.

- `dst` is a pointer to the target acceleration structure for the build.

- `src` is a pointer to an existing acceleration structure that is to be
  used to update the `dst` acceleration structure.

- `scratch` is the [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) that will be used as
  scratch memory for the build.

- `scratchOffset` is the offset in bytes relative to the start of
  `scratch` that will be used as a scratch memory.

## <a href="#_description" class="anchor"></a>Description

Accesses to `dst`, `src`, and `scratch` **must** be <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies"
target="_blank" rel="noopener">synchronized</a> with the
`VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
target="_blank" rel="noopener">pipeline stage</a> and an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-types"
target="_blank" rel="noopener">access type</a> of
`VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR` or
`VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_KHR`.

Valid Usage

- <a href="#VUID-vkCmdBuildAccelerationStructureNV-geometryCount-02241"
  id="VUID-vkCmdBuildAccelerationStructureNV-geometryCount-02241"></a>
  VUID-vkCmdBuildAccelerationStructureNV-geometryCount-02241  
  `geometryCount` **must** be less than or equal to
  [VkPhysicalDeviceRayTracingPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPropertiesNV.html)::`maxGeometryCount`

- <a href="#VUID-vkCmdBuildAccelerationStructureNV-dst-02488"
  id="VUID-vkCmdBuildAccelerationStructureNV-dst-02488"></a>
  VUID-vkCmdBuildAccelerationStructureNV-dst-02488  
  `dst` **must** have been created with compatible
  [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInfoNV.html)
  where
  [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInfoNV.html)::`type`
  and
  [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInfoNV.html)::`flags`
  are identical,
  [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInfoNV.html)::`instanceCount`
  and
  [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInfoNV.html)::`geometryCount`
  for `dst` are greater than or equal to the build size and each
  geometry in
  [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInfoNV.html)::`pGeometries`
  for `dst` has greater than or equal to the number of vertices,
  indices, and AABBs

- <a href="#VUID-vkCmdBuildAccelerationStructureNV-update-02489"
  id="VUID-vkCmdBuildAccelerationStructureNV-update-02489"></a>
  VUID-vkCmdBuildAccelerationStructureNV-update-02489  
  If `update` is `VK_TRUE`, `src` **must** not be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-vkCmdBuildAccelerationStructureNV-update-02490"
  id="VUID-vkCmdBuildAccelerationStructureNV-update-02490"></a>
  VUID-vkCmdBuildAccelerationStructureNV-update-02490  
  If `update` is `VK_TRUE`, `src` **must** have previously been
  constructed with `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_UPDATE_BIT_NV`
  set in
  [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInfoNV.html)::`flags`
  in the original build

- <a href="#VUID-vkCmdBuildAccelerationStructureNV-update-02491"
  id="VUID-vkCmdBuildAccelerationStructureNV-update-02491"></a>
  VUID-vkCmdBuildAccelerationStructureNV-update-02491  
  If `update` is `VK_FALSE`, the `size` member of the
  [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html) structure returned
  from a call to
  [vkGetAccelerationStructureMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureMemoryRequirementsNV.html)
  with
  [VkAccelerationStructureMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMemoryRequirementsInfoNV.html)::`accelerationStructure`
  set to `dst` and
  [VkAccelerationStructureMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMemoryRequirementsInfoNV.html)::`type`
  set to
  `VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_BUILD_SCRATCH_NV`
  **must** be less than or equal to the size of `scratch` minus
  `scratchOffset`

- <a href="#VUID-vkCmdBuildAccelerationStructureNV-update-02492"
  id="VUID-vkCmdBuildAccelerationStructureNV-update-02492"></a>
  VUID-vkCmdBuildAccelerationStructureNV-update-02492  
  If `update` is `VK_TRUE`, the `size` member of the
  [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html) structure returned
  from a call to
  [vkGetAccelerationStructureMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureMemoryRequirementsNV.html)
  with
  [VkAccelerationStructureMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMemoryRequirementsInfoNV.html)::`accelerationStructure`
  set to `dst` and
  [VkAccelerationStructureMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMemoryRequirementsInfoNV.html)::`type`
  set to
  `VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_UPDATE_SCRATCH_NV`
  **must** be less than or equal to the size of `scratch` minus
  `scratchOffset`

- <a href="#VUID-vkCmdBuildAccelerationStructureNV-scratch-03522"
  id="VUID-vkCmdBuildAccelerationStructureNV-scratch-03522"></a>
  VUID-vkCmdBuildAccelerationStructureNV-scratch-03522  
  `scratch` **must** have been created with
  `VK_BUFFER_USAGE_RAY_TRACING_BIT_NV` usage flag

- <a href="#VUID-vkCmdBuildAccelerationStructureNV-instanceData-03523"
  id="VUID-vkCmdBuildAccelerationStructureNV-instanceData-03523"></a>
  VUID-vkCmdBuildAccelerationStructureNV-instanceData-03523  
  If `instanceData` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `instanceData` **must** have been created with
  `VK_BUFFER_USAGE_RAY_TRACING_BIT_NV` usage flag

- <a
  href="#VUID-vkCmdBuildAccelerationStructureNV-accelerationStructureReference-03786"
  id="VUID-vkCmdBuildAccelerationStructureNV-accelerationStructureReference-03786"></a>
  VUID-vkCmdBuildAccelerationStructureNV-accelerationStructureReference-03786  
  Each
  [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInstanceKHR.html)::`accelerationStructureReference`
  value in `instanceData` **must** be a valid device address containing
  a value obtained from
  [vkGetAccelerationStructureHandleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureHandleNV.html)

- <a href="#VUID-vkCmdBuildAccelerationStructureNV-update-03524"
  id="VUID-vkCmdBuildAccelerationStructureNV-update-03524"></a>
  VUID-vkCmdBuildAccelerationStructureNV-update-03524  
  If `update` is `VK_TRUE`, then objects that were previously active
  **must** not be made inactive as per <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#acceleration-structure-inactive-prims"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#acceleration-structure-inactive-prims</a>

- <a href="#VUID-vkCmdBuildAccelerationStructureNV-update-03525"
  id="VUID-vkCmdBuildAccelerationStructureNV-update-03525"></a>
  VUID-vkCmdBuildAccelerationStructureNV-update-03525  
  If `update` is `VK_TRUE`, then objects that were previously inactive
  **must** not be made active as per <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#acceleration-structure-inactive-prims"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#acceleration-structure-inactive-prims</a>

- <a href="#VUID-vkCmdBuildAccelerationStructureNV-update-03526"
  id="VUID-vkCmdBuildAccelerationStructureNV-update-03526"></a>
  VUID-vkCmdBuildAccelerationStructureNV-update-03526  
  If `update` is `VK_TRUE`, the `src` and `dst` objects **must** either
  be the same object or not have any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-memory-aliasing"
  target="_blank" rel="noopener">memory aliasing</a>

- <a href="#VUID-vkCmdBuildAccelerationStructureNV-dst-07787"
  id="VUID-vkCmdBuildAccelerationStructureNV-dst-07787"></a>
  VUID-vkCmdBuildAccelerationStructureNV-dst-07787  
  `dst` **must** be bound completely and contiguously to a single
  `VkDeviceMemory` object via
  [vkBindAccelerationStructureMemoryNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindAccelerationStructureMemoryNV.html)

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdBuildAccelerationStructureNV-commandBuffer-parameter"
  id="VUID-vkCmdBuildAccelerationStructureNV-commandBuffer-parameter"></a>
  VUID-vkCmdBuildAccelerationStructureNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdBuildAccelerationStructureNV-pInfo-parameter"
  id="VUID-vkCmdBuildAccelerationStructureNV-pInfo-parameter"></a>
  VUID-vkCmdBuildAccelerationStructureNV-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInfoNV.html)
  structure

- <a href="#VUID-vkCmdBuildAccelerationStructureNV-instanceData-parameter"
  id="VUID-vkCmdBuildAccelerationStructureNV-instanceData-parameter"></a>
  VUID-vkCmdBuildAccelerationStructureNV-instanceData-parameter  
  If `instanceData` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `instanceData` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-vkCmdBuildAccelerationStructureNV-dst-parameter"
  id="VUID-vkCmdBuildAccelerationStructureNV-dst-parameter"></a>
  VUID-vkCmdBuildAccelerationStructureNV-dst-parameter  
  `dst` **must** be a valid
  [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html) handle

- <a href="#VUID-vkCmdBuildAccelerationStructureNV-src-parameter"
  id="VUID-vkCmdBuildAccelerationStructureNV-src-parameter"></a>
  VUID-vkCmdBuildAccelerationStructureNV-src-parameter  
  If `src` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `src` **must**
  be a valid [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html)
  handle

- <a href="#VUID-vkCmdBuildAccelerationStructureNV-scratch-parameter"
  id="VUID-vkCmdBuildAccelerationStructureNV-scratch-parameter"></a>
  VUID-vkCmdBuildAccelerationStructureNV-scratch-parameter  
  `scratch` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a
  href="#VUID-vkCmdBuildAccelerationStructureNV-commandBuffer-recording"
  id="VUID-vkCmdBuildAccelerationStructureNV-commandBuffer-recording"></a>
  VUID-vkCmdBuildAccelerationStructureNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdBuildAccelerationStructureNV-commandBuffer-cmdpool"
  id="VUID-vkCmdBuildAccelerationStructureNV-commandBuffer-cmdpool"></a>
  VUID-vkCmdBuildAccelerationStructureNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support compute operations

- <a href="#VUID-vkCmdBuildAccelerationStructureNV-renderpass"
  id="VUID-vkCmdBuildAccelerationStructureNV-renderpass"></a>
  VUID-vkCmdBuildAccelerationStructureNV-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdBuildAccelerationStructureNV-videocoding"
  id="VUID-vkCmdBuildAccelerationStructureNV-videocoding"></a>
  VUID-vkCmdBuildAccelerationStructureNV-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdBuildAccelerationStructureNV-commonparent"
  id="VUID-vkCmdBuildAccelerationStructureNV-commonparent"></a>
  VUID-vkCmdBuildAccelerationStructureNV-commonparent  
  Each of `commandBuffer`, `dst`, `instanceData`, `scratch`, and `src`
  that are valid handles of non-ignored parameters **must** have been
  created, allocated, or retrieved from the same
  [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

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

[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInfoNV.html),
[VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBuildAccelerationStructureNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
