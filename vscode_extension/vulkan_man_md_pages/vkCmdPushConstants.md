# vkCmdPushConstants(3) Manual Page

## Name

vkCmdPushConstants - Update the values of push constants



## <a href="#_c_specification" class="anchor"></a>C Specification

To update push constants, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdPushConstants(
    VkCommandBuffer                             commandBuffer,
    VkPipelineLayout                            layout,
    VkShaderStageFlags                          stageFlags,
    uint32_t                                    offset,
    uint32_t                                    size,
    const void*                                 pValues);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer in which the push constant
  update will be recorded.

- `layout` is the pipeline layout used to program the push constant
  updates.

- `stageFlags` is a bitmask of
  [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html) specifying the
  shader stages that will use the push constants in the updated range.

- `offset` is the start offset of the push constant range to update, in
  units of bytes.

- `size` is the size of the push constant range to update, in units of
  bytes.

- `pValues` is a pointer to an array of `size` bytes containing the new
  push constant values.

## <a href="#_description" class="anchor"></a>Description

When a command buffer begins recording, all push constant values are
undefined. Reads of undefined push constant values by the executing
shader return undefined values.

Push constant values **can** be updated incrementally, causing shader
stages in `stageFlags` to read the new data from `pValues` for push
constants modified by this command, while still reading the previous
data for push constants not modified by this command. When a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-bindpoint-commands"
target="_blank" rel="noopener">bound pipeline command</a> is issued, the
bound pipeline’s layout **must** be compatible with the layouts used to
set the values of all push constants in the pipeline layout’s push
constant ranges, as described in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-compatibility"
target="_blank" rel="noopener">Pipeline Layout Compatibility</a>.
Binding a pipeline with a layout that is not compatible with the push
constant layout does not disturb the push constant values.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>As <code>stageFlags</code> needs to include all flags the relevant
push constant ranges were created with, any flags that are not supported
by the queue family that the <a
href="VkCommandPool.html">VkCommandPool</a> used to allocate
<code>commandBuffer</code> was created on are ignored.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkCmdPushConstants-offset-01795"
  id="VUID-vkCmdPushConstants-offset-01795"></a>
  VUID-vkCmdPushConstants-offset-01795  
  For each byte in the range specified by `offset` and `size` and for
  each shader stage in `stageFlags`, there **must** be a push constant
  range in `layout` that includes that byte and that stage

- <a href="#VUID-vkCmdPushConstants-offset-01796"
  id="VUID-vkCmdPushConstants-offset-01796"></a>
  VUID-vkCmdPushConstants-offset-01796  
  For each byte in the range specified by `offset` and `size` and for
  each push constant range that overlaps that byte, `stageFlags`
  **must** include all stages in that push constant range’s
  [VkPushConstantRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPushConstantRange.html)::`stageFlags`

- <a href="#VUID-vkCmdPushConstants-offset-00368"
  id="VUID-vkCmdPushConstants-offset-00368"></a>
  VUID-vkCmdPushConstants-offset-00368  
  `offset` **must** be a multiple of `4`

- <a href="#VUID-vkCmdPushConstants-size-00369"
  id="VUID-vkCmdPushConstants-size-00369"></a>
  VUID-vkCmdPushConstants-size-00369  
  `size` **must** be a multiple of `4`

- <a href="#VUID-vkCmdPushConstants-offset-00370"
  id="VUID-vkCmdPushConstants-offset-00370"></a>
  VUID-vkCmdPushConstants-offset-00370  
  `offset` **must** be less than
  `VkPhysicalDeviceLimits`::`maxPushConstantsSize`

- <a href="#VUID-vkCmdPushConstants-size-00371"
  id="VUID-vkCmdPushConstants-size-00371"></a>
  VUID-vkCmdPushConstants-size-00371  
  `size` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxPushConstantsSize` minus `offset`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdPushConstants-commandBuffer-parameter"
  id="VUID-vkCmdPushConstants-commandBuffer-parameter"></a>
  VUID-vkCmdPushConstants-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdPushConstants-layout-parameter"
  id="VUID-vkCmdPushConstants-layout-parameter"></a>
  VUID-vkCmdPushConstants-layout-parameter  
  `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html)
  handle

- <a href="#VUID-vkCmdPushConstants-stageFlags-parameter"
  id="VUID-vkCmdPushConstants-stageFlags-parameter"></a>
  VUID-vkCmdPushConstants-stageFlags-parameter  
  `stageFlags` **must** be a valid combination of
  [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html) values

- <a href="#VUID-vkCmdPushConstants-stageFlags-requiredbitmask"
  id="VUID-vkCmdPushConstants-stageFlags-requiredbitmask"></a>
  VUID-vkCmdPushConstants-stageFlags-requiredbitmask  
  `stageFlags` **must** not be `0`

- <a href="#VUID-vkCmdPushConstants-pValues-parameter"
  id="VUID-vkCmdPushConstants-pValues-parameter"></a>
  VUID-vkCmdPushConstants-pValues-parameter  
  `pValues` **must** be a valid pointer to an array of `size` bytes

- <a href="#VUID-vkCmdPushConstants-commandBuffer-recording"
  id="VUID-vkCmdPushConstants-commandBuffer-recording"></a>
  VUID-vkCmdPushConstants-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdPushConstants-commandBuffer-cmdpool"
  id="VUID-vkCmdPushConstants-commandBuffer-cmdpool"></a>
  VUID-vkCmdPushConstants-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a href="#VUID-vkCmdPushConstants-videocoding"
  id="VUID-vkCmdPushConstants-videocoding"></a>
  VUID-vkCmdPushConstants-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdPushConstants-size-arraylength"
  id="VUID-vkCmdPushConstants-size-arraylength"></a>
  VUID-vkCmdPushConstants-size-arraylength  
  `size` **must** be greater than `0`

- <a href="#VUID-vkCmdPushConstants-commonparent"
  id="VUID-vkCmdPushConstants-commonparent"></a>
  VUID-vkCmdPushConstants-commonparent  
  Both of `commandBuffer`, and `layout` **must** have been created,
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
<td class="tableblock halign-left valign-top"><p>Both</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics<br />
Compute</p></td>
<td class="tableblock halign-left valign-top"><p>State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html),
[VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdPushConstants"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
