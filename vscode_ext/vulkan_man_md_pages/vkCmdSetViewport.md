# vkCmdSetViewport(3) Manual Page

## Name

vkCmdSetViewport - Set the viewport dynamically for a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the viewport
transformation parameters, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdSetViewport(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstViewport,
    uint32_t                                    viewportCount,
    const VkViewport*                           pViewports);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `firstViewport` is the index of the first viewport whose parameters
  are updated by the command.

- `viewportCount` is the number of viewports whose parameters are
  updated by the command.

- `pViewports` is a pointer to an array of [VkViewport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewport.html)
  structures specifying viewport parameters.

## <a href="#_description" class="anchor"></a>Description

This command sets the viewport transformation parameters state for
subsequent drawing commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_VIEWPORT` set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
`VkPipelineViewportStateCreateInfo`::`pViewports` values used to create
the currently active pipeline.

The viewport parameters taken from element i of `pViewports` replace the
current state for the viewport index `firstViewport` + i, for i in \[0,
`viewportCount`).

Valid Usage

- <a href="#VUID-vkCmdSetViewport-firstViewport-01223"
  id="VUID-vkCmdSetViewport-firstViewport-01223"></a>
  VUID-vkCmdSetViewport-firstViewport-01223  
  The sum of `firstViewport` and `viewportCount` **must** be between `1`
  and `VkPhysicalDeviceLimits`::`maxViewports`, inclusive

- <a href="#VUID-vkCmdSetViewport-firstViewport-01224"
  id="VUID-vkCmdSetViewport-firstViewport-01224"></a>
  VUID-vkCmdSetViewport-firstViewport-01224  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiViewport"
  target="_blank" rel="noopener"><code>multiViewport</code></a> feature
  is not enabled, `firstViewport` **must** be `0`

- <a href="#VUID-vkCmdSetViewport-viewportCount-01225"
  id="VUID-vkCmdSetViewport-viewportCount-01225"></a>
  VUID-vkCmdSetViewport-viewportCount-01225  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiViewport"
  target="_blank" rel="noopener"><code>multiViewport</code></a> feature
  is not enabled, `viewportCount` **must** be `1`

- <a href="#VUID-vkCmdSetViewport-commandBuffer-04821"
  id="VUID-vkCmdSetViewport-commandBuffer-04821"></a>
  VUID-vkCmdSetViewport-commandBuffer-04821  
  `commandBuffer` **must** not have
  [VkCommandBufferInheritanceViewportScissorInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceViewportScissorInfoNV.html)::`viewportScissor2D`
  enabled

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetViewport-commandBuffer-parameter"
  id="VUID-vkCmdSetViewport-commandBuffer-parameter"></a>
  VUID-vkCmdSetViewport-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetViewport-pViewports-parameter"
  id="VUID-vkCmdSetViewport-pViewports-parameter"></a>
  VUID-vkCmdSetViewport-pViewports-parameter  
  `pViewports` **must** be a valid pointer to an array of
  `viewportCount` valid [VkViewport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewport.html) structures

- <a href="#VUID-vkCmdSetViewport-commandBuffer-recording"
  id="VUID-vkCmdSetViewport-commandBuffer-recording"></a>
  VUID-vkCmdSetViewport-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetViewport-commandBuffer-cmdpool"
  id="VUID-vkCmdSetViewport-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetViewport-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetViewport-videocoding"
  id="VUID-vkCmdSetViewport-videocoding"></a>
  VUID-vkCmdSetViewport-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdSetViewport-viewportCount-arraylength"
  id="VUID-vkCmdSetViewport-viewportCount-arraylength"></a>
  VUID-vkCmdSetViewport-viewportCount-arraylength  
  `viewportCount` **must** be greater than `0`

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

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html), [VkViewport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewport.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetViewport"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
