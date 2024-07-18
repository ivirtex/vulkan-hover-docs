# vkCmdCopyAccelerationStructureNV(3) Manual Page

## Name

vkCmdCopyAccelerationStructureNV - Copy an acceleration structure



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy an acceleration structure call:

``` c
// Provided by VK_NV_ray_tracing
void vkCmdCopyAccelerationStructureNV(
    VkCommandBuffer                             commandBuffer,
    VkAccelerationStructureNV                   dst,
    VkAccelerationStructureNV                   src,
    VkCopyAccelerationStructureModeKHR          mode);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `dst` is the target acceleration structure for the copy.

- `src` is the source acceleration structure for the copy.

- `mode` is a
  [VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureModeKHR.html)
  value specifying additional operations to perform during the copy.

## <a href="#_description" class="anchor"></a>Description

Accesses to `src` and `dst` **must** be <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies"
target="_blank" rel="noopener">synchronized</a> with the
`VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR` <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
target="_blank" rel="noopener">pipeline stage</a> or the
`VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
target="_blank" rel="noopener">pipeline stage</a>, and an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-types"
target="_blank" rel="noopener">access type</a> of
`VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR` or
`VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_KHR` as appropriate.

Valid Usage

- <a href="#VUID-vkCmdCopyAccelerationStructureNV-mode-03410"
  id="VUID-vkCmdCopyAccelerationStructureNV-mode-03410"></a>
  VUID-vkCmdCopyAccelerationStructureNV-mode-03410  
  `mode` **must** be `VK_COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_KHR`
  or `VK_COPY_ACCELERATION_STRUCTURE_MODE_CLONE_KHR`

- <a href="#VUID-vkCmdCopyAccelerationStructureNV-src-04963"
  id="VUID-vkCmdCopyAccelerationStructureNV-src-04963"></a>
  VUID-vkCmdCopyAccelerationStructureNV-src-04963  
  The source acceleration structure `src` **must** have been constructed
  prior to the execution of this command

- <a href="#VUID-vkCmdCopyAccelerationStructureNV-src-03411"
  id="VUID-vkCmdCopyAccelerationStructureNV-src-03411"></a>
  VUID-vkCmdCopyAccelerationStructureNV-src-03411  
  If `mode` is `VK_COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_KHR`, `src`
  **must** have been constructed with
  `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_COMPACTION_BIT_KHR` in the
  build

- <a href="#VUID-vkCmdCopyAccelerationStructureNV-buffer-03718"
  id="VUID-vkCmdCopyAccelerationStructureNV-buffer-03718"></a>
  VUID-vkCmdCopyAccelerationStructureNV-buffer-03718  
  The `buffer` used to create `src` **must** be bound to device memory

- <a href="#VUID-vkCmdCopyAccelerationStructureNV-buffer-03719"
  id="VUID-vkCmdCopyAccelerationStructureNV-buffer-03719"></a>
  VUID-vkCmdCopyAccelerationStructureNV-buffer-03719  
  The `buffer` used to create `dst` **must** be bound to device memory

- <a href="#VUID-vkCmdCopyAccelerationStructureNV-dst-07791"
  id="VUID-vkCmdCopyAccelerationStructureNV-dst-07791"></a>
  VUID-vkCmdCopyAccelerationStructureNV-dst-07791  
  The range of memory backing `dst` that is accessed by this command
  **must** not overlap the memory backing `src` that is accessed by this
  command

- <a href="#VUID-vkCmdCopyAccelerationStructureNV-dst-07792"
  id="VUID-vkCmdCopyAccelerationStructureNV-dst-07792"></a>
  VUID-vkCmdCopyAccelerationStructureNV-dst-07792  
  `dst` **must** be bound completely and contiguously to a single
  `VkDeviceMemory` object via
  [vkBindAccelerationStructureMemoryNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindAccelerationStructureMemoryNV.html)

Valid Usage (Implicit)

- <a href="#VUID-vkCmdCopyAccelerationStructureNV-commandBuffer-parameter"
  id="VUID-vkCmdCopyAccelerationStructureNV-commandBuffer-parameter"></a>
  VUID-vkCmdCopyAccelerationStructureNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdCopyAccelerationStructureNV-dst-parameter"
  id="VUID-vkCmdCopyAccelerationStructureNV-dst-parameter"></a>
  VUID-vkCmdCopyAccelerationStructureNV-dst-parameter  
  `dst` **must** be a valid
  [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html) handle

- <a href="#VUID-vkCmdCopyAccelerationStructureNV-src-parameter"
  id="VUID-vkCmdCopyAccelerationStructureNV-src-parameter"></a>
  VUID-vkCmdCopyAccelerationStructureNV-src-parameter  
  `src` **must** be a valid
  [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html) handle

- <a href="#VUID-vkCmdCopyAccelerationStructureNV-mode-parameter"
  id="VUID-vkCmdCopyAccelerationStructureNV-mode-parameter"></a>
  VUID-vkCmdCopyAccelerationStructureNV-mode-parameter  
  `mode` **must** be a valid
  [VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureModeKHR.html)
  value

- <a href="#VUID-vkCmdCopyAccelerationStructureNV-commandBuffer-recording"
  id="VUID-vkCmdCopyAccelerationStructureNV-commandBuffer-recording"></a>
  VUID-vkCmdCopyAccelerationStructureNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdCopyAccelerationStructureNV-commandBuffer-cmdpool"
  id="VUID-vkCmdCopyAccelerationStructureNV-commandBuffer-cmdpool"></a>
  VUID-vkCmdCopyAccelerationStructureNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support compute operations

- <a href="#VUID-vkCmdCopyAccelerationStructureNV-renderpass"
  id="VUID-vkCmdCopyAccelerationStructureNV-renderpass"></a>
  VUID-vkCmdCopyAccelerationStructureNV-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdCopyAccelerationStructureNV-videocoding"
  id="VUID-vkCmdCopyAccelerationStructureNV-videocoding"></a>
  VUID-vkCmdCopyAccelerationStructureNV-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdCopyAccelerationStructureNV-commonparent"
  id="VUID-vkCmdCopyAccelerationStructureNV-commonparent"></a>
  VUID-vkCmdCopyAccelerationStructureNV-commonparent  
  Each of `commandBuffer`, `dst`, and `src` **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

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
[VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureModeKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdCopyAccelerationStructureNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
