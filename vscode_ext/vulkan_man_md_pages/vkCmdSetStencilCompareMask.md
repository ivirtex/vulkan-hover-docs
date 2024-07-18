# vkCmdSetStencilCompareMask(3) Manual Page

## Name

vkCmdSetStencilCompareMask - Set stencil compare mask dynamically for a
command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the stencil compare
mask, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdSetStencilCompareMask(
    VkCommandBuffer                             commandBuffer,
    VkStencilFaceFlags                          faceMask,
    uint32_t                                    compareMask);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `faceMask` is a bitmask of
  [VkStencilFaceFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilFaceFlagBits.html) specifying the set
  of stencil state for which to update the compare mask.

- `compareMask` is the new value to use as the stencil compare mask.

## <a href="#_description" class="anchor"></a>Description

This command sets the stencil compare mask for subsequent drawing
commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_STENCIL_COMPARE_MASK` set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkStencilOpState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOpState.html)::`compareMask` value used to
create the currently active pipeline, for both front and back faces.

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetStencilCompareMask-commandBuffer-parameter"
  id="VUID-vkCmdSetStencilCompareMask-commandBuffer-parameter"></a>
  VUID-vkCmdSetStencilCompareMask-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetStencilCompareMask-faceMask-parameter"
  id="VUID-vkCmdSetStencilCompareMask-faceMask-parameter"></a>
  VUID-vkCmdSetStencilCompareMask-faceMask-parameter  
  `faceMask` **must** be a valid combination of
  [VkStencilFaceFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilFaceFlagBits.html) values

- <a href="#VUID-vkCmdSetStencilCompareMask-faceMask-requiredbitmask"
  id="VUID-vkCmdSetStencilCompareMask-faceMask-requiredbitmask"></a>
  VUID-vkCmdSetStencilCompareMask-faceMask-requiredbitmask  
  `faceMask` **must** not be `0`

- <a href="#VUID-vkCmdSetStencilCompareMask-commandBuffer-recording"
  id="VUID-vkCmdSetStencilCompareMask-commandBuffer-recording"></a>
  VUID-vkCmdSetStencilCompareMask-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetStencilCompareMask-commandBuffer-cmdpool"
  id="VUID-vkCmdSetStencilCompareMask-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetStencilCompareMask-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetStencilCompareMask-videocoding"
  id="VUID-vkCmdSetStencilCompareMask-videocoding"></a>
  VUID-vkCmdSetStencilCompareMask-videocoding  
  This command **must** only be called outside of a video coding scope

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
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkStencilFaceFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilFaceFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetStencilCompareMask"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
