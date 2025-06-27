# vkCmdSetPerformanceOverrideINTEL(3) Manual Page

## Name

vkCmdSetPerformanceOverrideINTEL - Performance override settings



## <a href="#_c_specification" class="anchor"></a>C Specification

Some applications might want measure the effect of a set of commands
with a different settings. It is possible to override a particular
settings using :

``` c
// Provided by VK_INTEL_performance_query
VkResult vkCmdSetPerformanceOverrideINTEL(
    VkCommandBuffer                             commandBuffer,
    const VkPerformanceOverrideInfoINTEL*       pOverrideInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer where the override takes place.

- `pOverrideInfo` is a pointer to a
  [VkPerformanceOverrideInfoINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceOverrideInfoINTEL.html)
  structure selecting the parameter to override.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkCmdSetPerformanceOverrideINTEL-pOverrideInfo-02736"
  id="VUID-vkCmdSetPerformanceOverrideINTEL-pOverrideInfo-02736"></a>
  VUID-vkCmdSetPerformanceOverrideINTEL-pOverrideInfo-02736  
  `pOverrideInfo` **must** not be used with a
  [VkPerformanceOverrideTypeINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceOverrideTypeINTEL.html)
  that is not reported available by `vkGetPerformanceParameterINTEL`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetPerformanceOverrideINTEL-commandBuffer-parameter"
  id="VUID-vkCmdSetPerformanceOverrideINTEL-commandBuffer-parameter"></a>
  VUID-vkCmdSetPerformanceOverrideINTEL-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetPerformanceOverrideINTEL-pOverrideInfo-parameter"
  id="VUID-vkCmdSetPerformanceOverrideINTEL-pOverrideInfo-parameter"></a>
  VUID-vkCmdSetPerformanceOverrideINTEL-pOverrideInfo-parameter  
  `pOverrideInfo` **must** be a valid pointer to a valid
  [VkPerformanceOverrideInfoINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceOverrideInfoINTEL.html)
  structure

- <a href="#VUID-vkCmdSetPerformanceOverrideINTEL-commandBuffer-recording"
  id="VUID-vkCmdSetPerformanceOverrideINTEL-commandBuffer-recording"></a>
  VUID-vkCmdSetPerformanceOverrideINTEL-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetPerformanceOverrideINTEL-commandBuffer-cmdpool"
  id="VUID-vkCmdSetPerformanceOverrideINTEL-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetPerformanceOverrideINTEL-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, compute, or transfer operations

- <a href="#VUID-vkCmdSetPerformanceOverrideINTEL-videocoding"
  id="VUID-vkCmdSetPerformanceOverrideINTEL-videocoding"></a>
  VUID-vkCmdSetPerformanceOverrideINTEL-videocoding  
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
<td class="tableblock halign-left valign-top"><p>Graphics<br />
Compute<br />
Transfer</p></td>
<td class="tableblock halign-left valign-top"><p>State</p></td>
</tr>
</tbody>
</table>

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_TOO_MANY_OBJECTS`

- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_INTEL_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_INTEL_performance_query.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkPerformanceOverrideInfoINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceOverrideInfoINTEL.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetPerformanceOverrideINTEL"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
