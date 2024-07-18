# vkCmdSetDeviceMask(3) Manual Page

## Name

vkCmdSetDeviceMask - Modify device mask of a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To update the current device mask of a command buffer, call:

``` c
// Provided by VK_VERSION_1_1
void vkCmdSetDeviceMask(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    deviceMask);
```

or the equivalent command

``` c
// Provided by VK_KHR_device_group
void vkCmdSetDeviceMaskKHR(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    deviceMask);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is command buffer whose current device mask is
  modified.

- `deviceMask` is the new value of the current device mask.

## <a href="#_description" class="anchor"></a>Description

`deviceMask` is used to filter out subsequent commands from executing on
all physical devices whose bit indices are not set in the mask, except
commands beginning a render pass instance, commands transitioning to the
next subpass in the render pass instance, and commands ending a render
pass instance, which always execute on the set of physical devices whose
bit indices are included in the `deviceMask` member of the
[VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)
structure passed to the command beginning the corresponding render pass
instance.

Valid Usage

- <a href="#VUID-vkCmdSetDeviceMask-deviceMask-00108"
  id="VUID-vkCmdSetDeviceMask-deviceMask-00108"></a>
  VUID-vkCmdSetDeviceMask-deviceMask-00108  
  `deviceMask` **must** be a valid device mask value

- <a href="#VUID-vkCmdSetDeviceMask-deviceMask-00109"
  id="VUID-vkCmdSetDeviceMask-deviceMask-00109"></a>
  VUID-vkCmdSetDeviceMask-deviceMask-00109  
  `deviceMask` **must** not be zero

- <a href="#VUID-vkCmdSetDeviceMask-deviceMask-00110"
  id="VUID-vkCmdSetDeviceMask-deviceMask-00110"></a>
  VUID-vkCmdSetDeviceMask-deviceMask-00110  
  `deviceMask` **must** not include any set bits that were not in the
  [VkDeviceGroupCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupCommandBufferBeginInfo.html)::`deviceMask`
  value when the command buffer began recording

- <a href="#VUID-vkCmdSetDeviceMask-deviceMask-00111"
  id="VUID-vkCmdSetDeviceMask-deviceMask-00111"></a>
  VUID-vkCmdSetDeviceMask-deviceMask-00111  
  If `vkCmdSetDeviceMask` is called inside a render pass instance,
  `deviceMask` **must** not include any set bits that were not in the
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)::`deviceMask`
  value when the render pass instance began recording

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetDeviceMask-commandBuffer-parameter"
  id="VUID-vkCmdSetDeviceMask-commandBuffer-parameter"></a>
  VUID-vkCmdSetDeviceMask-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetDeviceMask-commandBuffer-recording"
  id="VUID-vkCmdSetDeviceMask-commandBuffer-recording"></a>
  VUID-vkCmdSetDeviceMask-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetDeviceMask-commandBuffer-cmdpool"
  id="VUID-vkCmdSetDeviceMask-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetDeviceMask-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, compute, or transfer operations

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
<td class="tableblock halign-left valign-top"><p>Both</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics<br />
Compute<br />
Transfer</p></td>
<td class="tableblock halign-left valign-top"><p>State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetDeviceMask"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
