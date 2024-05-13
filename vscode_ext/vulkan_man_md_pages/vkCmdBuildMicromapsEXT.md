# vkCmdBuildMicromapsEXT(3) Manual Page

## Name

vkCmdBuildMicromapsEXT - Build a micromap



## <a href="#_c_specification" class="anchor"></a>C Specification

To build micromaps call:

``` c
// Provided by VK_EXT_opacity_micromap
void vkCmdBuildMicromapsEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    infoCount,
    const VkMicromapBuildInfoEXT*               pInfos);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `infoCount` is the number of micromaps to build. It specifies the
  number of the `pInfos` structures that **must** be provided.

- `pInfos` is a pointer to an array of `infoCount`
  [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapBuildInfoEXT.html) structures
  defining the data used to build each micromap.

## <a href="#_description" class="anchor"></a>Description

The `vkCmdBuildMicromapsEXT` command provides the ability to initiate
multiple micromaps builds, however there is no ordering or
synchronization implied between any of the individual micromap builds.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>This means that there <strong>cannot</strong> be any memory aliasing
between any micromap memories or scratch memories being used by any of
the builds.</p></td>
</tr>
</tbody>
</table>

Accesses to the micromap scratch buffers as identified by the
[VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapBuildInfoEXT.html)::`scratchData`
buffer device addresses **must** be <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies"
target="_blank" rel="noopener">synchronized</a> with the
`VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
target="_blank" rel="noopener">pipeline stage</a> and an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-types"
target="_blank" rel="noopener">access type</a> of
(`VK_ACCESS_2_MICROMAP_READ_BIT_EXT` \|
`VK_ACCESS_2_MICROMAP_WRITE_BIT_EXT`). Accesses to
[VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapBuildInfoEXT.html)::`dstMicromap`
**must** be <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies"
target="_blank" rel="noopener">synchronized</a> with the
`VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
target="_blank" rel="noopener">pipeline stage</a> and an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-types"
target="_blank" rel="noopener">access type</a> of
`VK_ACCESS_2_MICROMAP_WRITE_BIT_EXT`.

Accesses to other input buffers as identified by any used values of
[VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapBuildInfoEXT.html)::`data` or
[VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapBuildInfoEXT.html)::`triangleArray`
**must** be <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies"
target="_blank" rel="noopener">synchronized</a> with the
`VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
target="_blank" rel="noopener">pipeline stage</a> and an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-types"
target="_blank" rel="noopener">access type</a> of
`VK_ACCESS_SHADER_READ_BIT`.

Valid Usage

- <a href="#VUID-vkCmdBuildMicromapsEXT-pInfos-07461"
  id="VUID-vkCmdBuildMicromapsEXT-pInfos-07461"></a>
  VUID-vkCmdBuildMicromapsEXT-pInfos-07461  
  For each `pInfos`\[i\], `dstMicromap` **must** have been created with
  a value of
  [VkMicromapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapCreateInfoEXT.html)::`size`
  greater than or equal to the memory size required by the build
  operation, as returned by
  [vkGetMicromapBuildSizesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMicromapBuildSizesEXT.html) with
  `pBuildInfo` = `pInfos`\[i\]

- <a href="#VUID-vkCmdBuildMicromapsEXT-mode-07462"
  id="VUID-vkCmdBuildMicromapsEXT-mode-07462"></a>
  VUID-vkCmdBuildMicromapsEXT-mode-07462  
  The `mode` member of each element of `pInfos` **must** be a valid
  [VkBuildMicromapModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildMicromapModeEXT.html) value

- <a href="#VUID-vkCmdBuildMicromapsEXT-dstMicromap-07463"
  id="VUID-vkCmdBuildMicromapsEXT-dstMicromap-07463"></a>
  VUID-vkCmdBuildMicromapsEXT-dstMicromap-07463  
  The `dstMicromap` member of any element of `pInfos` **must** be a
  valid [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapEXT.html) handle

- <a href="#VUID-vkCmdBuildMicromapsEXT-pInfos-07464"
  id="VUID-vkCmdBuildMicromapsEXT-pInfos-07464"></a>
  VUID-vkCmdBuildMicromapsEXT-pInfos-07464  
  For each element of `pInfos` its `type` member **must** match the
  value of
  [VkMicromapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapCreateInfoEXT.html)::`type` when
  its `dstMicromap` was created

- <a href="#VUID-vkCmdBuildMicromapsEXT-dstMicromap-07465"
  id="VUID-vkCmdBuildMicromapsEXT-dstMicromap-07465"></a>
  VUID-vkCmdBuildMicromapsEXT-dstMicromap-07465  
  The range of memory backing the `dstMicromap` member of any element of
  `pInfos` that is accessed by this command **must** not overlap the
  memory backing the `dstMicromap` member of any other element of
  `pInfos`, which is accessed by this command

- <a href="#VUID-vkCmdBuildMicromapsEXT-dstMicromap-07466"
  id="VUID-vkCmdBuildMicromapsEXT-dstMicromap-07466"></a>
  VUID-vkCmdBuildMicromapsEXT-dstMicromap-07466  
  The range of memory backing the `dstMicromap` member of any element of
  `pInfos` that is accessed by this command **must** not overlap the
  memory backing the `scratchData` member of any element of `pInfos`
  (including the same element), which is accessed by this command

- <a href="#VUID-vkCmdBuildMicromapsEXT-scratchData-07467"
  id="VUID-vkCmdBuildMicromapsEXT-scratchData-07467"></a>
  VUID-vkCmdBuildMicromapsEXT-scratchData-07467  
  The range of memory backing the `scratchData` member of any element of
  `pInfos` that is accessed by this command **must** not overlap the
  memory backing the `scratchData` member of any other element of
  `pInfos`, which is accessed by this command

<!-- -->

- <a href="#VUID-vkCmdBuildMicromapsEXT-pInfos-07508"
  id="VUID-vkCmdBuildMicromapsEXT-pInfos-07508"></a>
  VUID-vkCmdBuildMicromapsEXT-pInfos-07508  
  For each element of `pInfos`, the `buffer` used to create its
  `dstMicromap` member **must** be bound to device memory

- <a href="#VUID-vkCmdBuildMicromapsEXT-pInfos-07509"
  id="VUID-vkCmdBuildMicromapsEXT-pInfos-07509"></a>
  VUID-vkCmdBuildMicromapsEXT-pInfos-07509  
  If `pInfos`\[i\].`mode` is `VK_BUILD_MICROMAP_MODE_BUILD_EXT`, all
  addresses between `pInfos`\[i\].`scratchData.deviceAddress` and
  `pInfos`\[i\].`scratchData.deviceAddress` + N - 1 **must** be in the
  buffer device address range of the same buffer, where N is given by
  the `buildScratchSize` member of the
  [VkMicromapBuildSizesInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapBuildSizesInfoEXT.html)
  structure returned from a call to
  [vkGetMicromapBuildSizesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMicromapBuildSizesEXT.html) with an
  identical [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapBuildInfoEXT.html)
  structure and primitive count

- <a href="#VUID-vkCmdBuildMicromapsEXT-data-07510"
  id="VUID-vkCmdBuildMicromapsEXT-data-07510"></a>
  VUID-vkCmdBuildMicromapsEXT-data-07510  
  The buffers from which the buffer device addresses for all of the
  `data` and `triangleArray` members of all `pInfos`\[i\] are queried
  **must** have been created with the
  `VK_BUFFER_USAGE_MICROMAP_BUILD_INPUT_READ_ONLY_BIT_EXT` usage flag

- <a href="#VUID-vkCmdBuildMicromapsEXT-pInfos-07511"
  id="VUID-vkCmdBuildMicromapsEXT-pInfos-07511"></a>
  VUID-vkCmdBuildMicromapsEXT-pInfos-07511  
  For each element of `pInfos`\[i\] the buffer from which the buffer
  device address `pInfos`\[i\].`scratchData.deviceAddress` is queried
  **must** have been created with `VK_BUFFER_USAGE_STORAGE_BUFFER_BIT`
  usage flag

- <a href="#VUID-vkCmdBuildMicromapsEXT-pInfos-07512"
  id="VUID-vkCmdBuildMicromapsEXT-pInfos-07512"></a>
  VUID-vkCmdBuildMicromapsEXT-pInfos-07512  
  For each element of `pInfos`, its `scratchData.deviceAddress`,
  `data.deviceAddress`, and `triangleArray.deviceAddress` members
  **must** be valid device addresses obtained from
  [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html)

- <a href="#VUID-vkCmdBuildMicromapsEXT-pInfos-07513"
  id="VUID-vkCmdBuildMicromapsEXT-pInfos-07513"></a>
  VUID-vkCmdBuildMicromapsEXT-pInfos-07513  
  For each element of `pInfos`, if `scratchData.deviceAddress`,
  `data.deviceAddress`, or `triangleArray.deviceAddress` is the address
  of a non-sparse buffer then it **must** be bound completely and
  contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object

- <a href="#VUID-vkCmdBuildMicromapsEXT-pInfos-07514"
  id="VUID-vkCmdBuildMicromapsEXT-pInfos-07514"></a>
  VUID-vkCmdBuildMicromapsEXT-pInfos-07514  
  For each element of `pInfos`, its `scratchData.deviceAddress` member
  **must** be a multiple of
  [VkPhysicalDeviceAccelerationStructurePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceAccelerationStructurePropertiesKHR.html)::`minAccelerationStructureScratchOffsetAlignment`

- <a href="#VUID-vkCmdBuildMicromapsEXT-pInfos-07515"
  id="VUID-vkCmdBuildMicromapsEXT-pInfos-07515"></a>
  VUID-vkCmdBuildMicromapsEXT-pInfos-07515  
  For each element of `pInfos`, its `triangleArray.deviceAddress` and
  `data.deviceAddress` members **must** be a multiple of `256`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdBuildMicromapsEXT-commandBuffer-parameter"
  id="VUID-vkCmdBuildMicromapsEXT-commandBuffer-parameter"></a>
  VUID-vkCmdBuildMicromapsEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdBuildMicromapsEXT-pInfos-parameter"
  id="VUID-vkCmdBuildMicromapsEXT-pInfos-parameter"></a>
  VUID-vkCmdBuildMicromapsEXT-pInfos-parameter  
  `pInfos` **must** be a valid pointer to an array of `infoCount` valid
  [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapBuildInfoEXT.html) structures

- <a href="#VUID-vkCmdBuildMicromapsEXT-commandBuffer-recording"
  id="VUID-vkCmdBuildMicromapsEXT-commandBuffer-recording"></a>
  VUID-vkCmdBuildMicromapsEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdBuildMicromapsEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdBuildMicromapsEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdBuildMicromapsEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support compute operations

- <a href="#VUID-vkCmdBuildMicromapsEXT-renderpass"
  id="VUID-vkCmdBuildMicromapsEXT-renderpass"></a>
  VUID-vkCmdBuildMicromapsEXT-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdBuildMicromapsEXT-videocoding"
  id="VUID-vkCmdBuildMicromapsEXT-videocoding"></a>
  VUID-vkCmdBuildMicromapsEXT-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdBuildMicromapsEXT-infoCount-arraylength"
  id="VUID-vkCmdBuildMicromapsEXT-infoCount-arraylength"></a>
  VUID-vkCmdBuildMicromapsEXT-infoCount-arraylength  
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
<tr class="header">
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
<tr class="odd">
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

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapBuildInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBuildMicromapsEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
