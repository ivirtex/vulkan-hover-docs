# vkCmdSetExclusiveScissorEnableNV(3) Manual Page

## Name

vkCmdSetExclusiveScissorEnableNV - Dynamically enable each exclusive
scissor for a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> whether an exclusive
scissor is enabled or not, call:

``` c
// Provided by VK_NV_scissor_exclusive
void vkCmdSetExclusiveScissorEnableNV(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstExclusiveScissor,
    uint32_t                                    exclusiveScissorCount,
    const VkBool32*                             pExclusiveScissorEnables);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `firstExclusiveScissor` is the index of the first exclusive scissor
  rectangle whose state is updated by the command.

- `exclusiveScissorCount` is the number of exclusive scissor rectangles
  updated by the command.

- `pExclusiveScissorEnables` is a pointer to an array of
  [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html) values defining whether the exclusive
  scissor is enabled.

## <a href="#_description" class="anchor"></a>Description

The exclusive scissor enables taken from element i of
`pExclusiveScissorEnables` replace the current state for the scissor
index `firstExclusiveScissor` + i, for i in \[0,
`exclusiveScissorCount`).

This command sets the exclusive scissor enable for subsequent drawing
commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_EXCLUSIVE_SCISSOR_ENABLE_NV`
set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is implied by the
[VkPipelineViewportExclusiveScissorStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportExclusiveScissorStateCreateInfoNV.html)::`exclusiveScissorCount`
value used to create the currently active pipeline, where all
`exclusiveScissorCount` exclusive scissors are implicitly enabled and
the remainder up to `VkPhysicalDeviceLimits`::`maxViewports` are
implicitly disabled.

Valid Usage

- <a href="#VUID-vkCmdSetExclusiveScissorEnableNV-exclusiveScissor-07853"
  id="VUID-vkCmdSetExclusiveScissorEnableNV-exclusiveScissor-07853"></a>
  VUID-vkCmdSetExclusiveScissorEnableNV-exclusiveScissor-07853  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-exclusiveScissor"
  target="_blank" rel="noopener"><code>exclusiveScissor</code></a>
  feature **must** be enabled, and the implementation **must** support
  at least `specVersion` `2` of the
  [`VK_NV_scissor_exclusive`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_scissor_exclusive.html) extension

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetExclusiveScissorEnableNV-commandBuffer-parameter"
  id="VUID-vkCmdSetExclusiveScissorEnableNV-commandBuffer-parameter"></a>
  VUID-vkCmdSetExclusiveScissorEnableNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdSetExclusiveScissorEnableNV-pExclusiveScissorEnables-parameter"
  id="VUID-vkCmdSetExclusiveScissorEnableNV-pExclusiveScissorEnables-parameter"></a>
  VUID-vkCmdSetExclusiveScissorEnableNV-pExclusiveScissorEnables-parameter  
  `pExclusiveScissorEnables` **must** be a valid pointer to an array of
  `exclusiveScissorCount` [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html) values

- <a href="#VUID-vkCmdSetExclusiveScissorEnableNV-commandBuffer-recording"
  id="VUID-vkCmdSetExclusiveScissorEnableNV-commandBuffer-recording"></a>
  VUID-vkCmdSetExclusiveScissorEnableNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetExclusiveScissorEnableNV-commandBuffer-cmdpool"
  id="VUID-vkCmdSetExclusiveScissorEnableNV-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetExclusiveScissorEnableNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetExclusiveScissorEnableNV-videocoding"
  id="VUID-vkCmdSetExclusiveScissorEnableNV-videocoding"></a>
  VUID-vkCmdSetExclusiveScissorEnableNV-videocoding  
  This command **must** only be called outside of a video coding scope

- <a
  href="#VUID-vkCmdSetExclusiveScissorEnableNV-exclusiveScissorCount-arraylength"
  id="VUID-vkCmdSetExclusiveScissorEnableNV-exclusiveScissorCount-arraylength"></a>
  VUID-vkCmdSetExclusiveScissorEnableNV-exclusiveScissorCount-arraylength  
  `exclusiveScissorCount` **must** be greater than `0`

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
<td class="tableblock halign-left valign-top"><p>Graphics</p></td>
<td class="tableblock halign-left valign-top"><p>State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_scissor_exclusive](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_scissor_exclusive.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetExclusiveScissorEnableNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
