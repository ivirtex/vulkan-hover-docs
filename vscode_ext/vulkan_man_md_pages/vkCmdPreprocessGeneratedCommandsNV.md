# vkCmdPreprocessGeneratedCommandsNV(3) Manual Page

## Name

vkCmdPreprocessGeneratedCommandsNV - Performs preprocessing for
generated commands



## <a href="#_c_specification" class="anchor"></a>C Specification

Commands **can** be preprocessed prior execution using the following
command:

``` c
// Provided by VK_NV_device_generated_commands
void vkCmdPreprocessGeneratedCommandsNV(
    VkCommandBuffer                             commandBuffer,
    const VkGeneratedCommandsInfoNV*            pGeneratedCommandsInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer which does the preprocessing.

- `pGeneratedCommandsInfo` is a pointer to a
  [VkGeneratedCommandsInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeneratedCommandsInfoNV.html) structure
  containing parameters affecting the preprocessing step.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkCmdPreprocessGeneratedCommandsNV-commandBuffer-02974"
  id="VUID-vkCmdPreprocessGeneratedCommandsNV-commandBuffer-02974"></a>
  VUID-vkCmdPreprocessGeneratedCommandsNV-commandBuffer-02974  
  `commandBuffer` **must** not be a protected command buffer

- <a
  href="#VUID-vkCmdPreprocessGeneratedCommandsNV-pGeneratedCommandsInfo-02927"
  id="VUID-vkCmdPreprocessGeneratedCommandsNV-pGeneratedCommandsInfo-02927"></a>
  VUID-vkCmdPreprocessGeneratedCommandsNV-pGeneratedCommandsInfo-02927  
  `pGeneratedCommandsInfo`\`s `indirectCommandsLayout` **must** have
  been created with the
  `VK_INDIRECT_COMMANDS_LAYOUT_USAGE_EXPLICIT_PREPROCESS_BIT_NV` bit set

- <a
  href="#VUID-vkCmdPreprocessGeneratedCommandsNV-deviceGeneratedCommands-02928"
  id="VUID-vkCmdPreprocessGeneratedCommandsNV-deviceGeneratedCommands-02928"></a>
  VUID-vkCmdPreprocessGeneratedCommandsNV-deviceGeneratedCommands-02928  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-deviceGeneratedCommands"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceDeviceGeneratedCommandsFeaturesNV</code>::<code>deviceGeneratedCommands</code></a>
  feature **must** be enabled

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdPreprocessGeneratedCommandsNV-commandBuffer-parameter"
  id="VUID-vkCmdPreprocessGeneratedCommandsNV-commandBuffer-parameter"></a>
  VUID-vkCmdPreprocessGeneratedCommandsNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdPreprocessGeneratedCommandsNV-pGeneratedCommandsInfo-parameter"
  id="VUID-vkCmdPreprocessGeneratedCommandsNV-pGeneratedCommandsInfo-parameter"></a>
  VUID-vkCmdPreprocessGeneratedCommandsNV-pGeneratedCommandsInfo-parameter  
  `pGeneratedCommandsInfo` **must** be a valid pointer to a valid
  [VkGeneratedCommandsInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeneratedCommandsInfoNV.html) structure

- <a
  href="#VUID-vkCmdPreprocessGeneratedCommandsNV-commandBuffer-recording"
  id="VUID-vkCmdPreprocessGeneratedCommandsNV-commandBuffer-recording"></a>
  VUID-vkCmdPreprocessGeneratedCommandsNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdPreprocessGeneratedCommandsNV-commandBuffer-cmdpool"
  id="VUID-vkCmdPreprocessGeneratedCommandsNV-commandBuffer-cmdpool"></a>
  VUID-vkCmdPreprocessGeneratedCommandsNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a href="#VUID-vkCmdPreprocessGeneratedCommandsNV-renderpass"
  id="VUID-vkCmdPreprocessGeneratedCommandsNV-renderpass"></a>
  VUID-vkCmdPreprocessGeneratedCommandsNV-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdPreprocessGeneratedCommandsNV-videocoding"
  id="VUID-vkCmdPreprocessGeneratedCommandsNV-videocoding"></a>
  VUID-vkCmdPreprocessGeneratedCommandsNV-videocoding  
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
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics<br />
Compute</p></td>
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkGeneratedCommandsInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeneratedCommandsInfoNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdPreprocessGeneratedCommandsNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
