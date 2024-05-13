# vkCmdSetColorWriteEnableEXT(3) Manual Page

## Name

vkCmdSetColorWriteEnableEXT - Enable or disable writes to a color
attachment dynamically for a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically enable or disable</a> writes
to a color attachment, call:

``` c
// Provided by VK_EXT_color_write_enable
void                                    vkCmdSetColorWriteEnableEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    attachmentCount,
    const VkBool32*                             pColorWriteEnables);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `attachmentCount` is the number of [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html) elements
  in `pColorWriteEnables`.

- `pColorWriteEnables` is a pointer to an array of per target attachment
  boolean values specifying whether color writes are enabled for the
  given attachment.

## <a href="#_description" class="anchor"></a>Description

This command sets the color write enables for subsequent drawing
commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_COLOR_WRITE_ENABLE_EXT` set
in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineColorWriteCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorWriteCreateInfoEXT.html)::`pColorWriteEnables`
values used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetColorWriteEnableEXT-None-04803"
  id="VUID-vkCmdSetColorWriteEnableEXT-None-04803"></a>
  VUID-vkCmdSetColorWriteEnableEXT-None-04803  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-colorWriteEnable"
  target="_blank" rel="noopener"><code>colorWriteEnable</code></a>
  feature **must** be enabled

- <a href="#VUID-vkCmdSetColorWriteEnableEXT-attachmentCount-06656"
  id="VUID-vkCmdSetColorWriteEnableEXT-attachmentCount-06656"></a>
  VUID-vkCmdSetColorWriteEnableEXT-attachmentCount-06656  
  `attachmentCount` **must** be less than or equal to the
  `maxColorAttachments` member of `VkPhysicalDeviceLimits`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetColorWriteEnableEXT-commandBuffer-parameter"
  id="VUID-vkCmdSetColorWriteEnableEXT-commandBuffer-parameter"></a>
  VUID-vkCmdSetColorWriteEnableEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetColorWriteEnableEXT-pColorWriteEnables-parameter"
  id="VUID-vkCmdSetColorWriteEnableEXT-pColorWriteEnables-parameter"></a>
  VUID-vkCmdSetColorWriteEnableEXT-pColorWriteEnables-parameter  
  `pColorWriteEnables` **must** be a valid pointer to an array of
  `attachmentCount` [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html) values

- <a href="#VUID-vkCmdSetColorWriteEnableEXT-commandBuffer-recording"
  id="VUID-vkCmdSetColorWriteEnableEXT-commandBuffer-recording"></a>
  VUID-vkCmdSetColorWriteEnableEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetColorWriteEnableEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdSetColorWriteEnableEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetColorWriteEnableEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetColorWriteEnableEXT-videocoding"
  id="VUID-vkCmdSetColorWriteEnableEXT-videocoding"></a>
  VUID-vkCmdSetColorWriteEnableEXT-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdSetColorWriteEnableEXT-attachmentCount-arraylength"
  id="VUID-vkCmdSetColorWriteEnableEXT-attachmentCount-arraylength"></a>
  VUID-vkCmdSetColorWriteEnableEXT-attachmentCount-arraylength  
  `attachmentCount` **must** be greater than `0`

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
<td class="tableblock halign-left valign-top"><p>Both</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics</p></td>
<td class="tableblock halign-left valign-top"><p>State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_color_write_enable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_color_write_enable.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetColorWriteEnableEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
