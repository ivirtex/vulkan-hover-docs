# vkEndCommandBuffer(3) Manual Page

## Name

vkEndCommandBuffer - Finish recording a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To complete recording of a command buffer, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkEndCommandBuffer(
    VkCommandBuffer                             commandBuffer);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer to complete recording.

## <a href="#_description" class="anchor"></a>Description

The command buffer **must** have been in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
target="_blank" rel="noopener">recording state</a>, and, if successful,
is moved to the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
target="_blank" rel="noopener">executable state</a>.

If there was an error during recording, the application will be notified
by an unsuccessful return code returned by `vkEndCommandBuffer`, and the
command buffer will be moved to the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
target="_blank" rel="noopener">invalid state</a>.

In case the application recorded one or more <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-encode-operations"
target="_blank" rel="noopener">video encode operations</a> into the
command buffer, implementations **may** return the
`VK_ERROR_INVALID_VIDEO_STD_PARAMETERS_KHR` error if any of the
specified Video Std parameters do not adhere to the syntactic or
semantic requirements of the used video compression standard, or if
values derived from parameters according to the rules defined by the
used video compression standard do not adhere to the capabilities of the
video compression standard or the implementation.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Applications <strong>should</strong> not rely on the
<code>VK_ERROR_INVALID_VIDEO_STD_PARAMETERS_KHR</code> error being
returned by any command as a means to verify Video Std parameters, as
implementations are not required to report the error in any specific set
of cases.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkEndCommandBuffer-commandBuffer-00059"
  id="VUID-vkEndCommandBuffer-commandBuffer-00059"></a>
  VUID-vkEndCommandBuffer-commandBuffer-00059  
  `commandBuffer` **must** be in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
  target="_blank" rel="noopener">recording state</a>

- <a href="#VUID-vkEndCommandBuffer-commandBuffer-00060"
  id="VUID-vkEndCommandBuffer-commandBuffer-00060"></a>
  VUID-vkEndCommandBuffer-commandBuffer-00060  
  If `commandBuffer` is a primary command buffer, there **must** not be
  an active render pass instance

- <a href="#VUID-vkEndCommandBuffer-commandBuffer-00061"
  id="VUID-vkEndCommandBuffer-commandBuffer-00061"></a>
  VUID-vkEndCommandBuffer-commandBuffer-00061  
  All queries made <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-operation-active"
  target="_blank" rel="noopener">active</a> during the recording of
  `commandBuffer` **must** have been made inactive

- <a href="#VUID-vkEndCommandBuffer-None-01978"
  id="VUID-vkEndCommandBuffer-None-01978"></a>
  VUID-vkEndCommandBuffer-None-01978  
  Conditional rendering **must** not be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#active-conditional-rendering"
  target="_blank" rel="noopener">active</a>

- <a href="#VUID-vkEndCommandBuffer-None-06991"
  id="VUID-vkEndCommandBuffer-None-06991"></a>
  VUID-vkEndCommandBuffer-None-06991  
  There **must** be no video session object bound

- <a href="#VUID-vkEndCommandBuffer-commandBuffer-01815"
  id="VUID-vkEndCommandBuffer-commandBuffer-01815"></a>
  VUID-vkEndCommandBuffer-commandBuffer-01815  
  If `commandBuffer` is a secondary command buffer, there **must** not
  be an outstanding
  [vkCmdBeginDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginDebugUtilsLabelEXT.html)
  command recorded to `commandBuffer` that has not previously been ended
  by a call to
  [vkCmdEndDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndDebugUtilsLabelEXT.html)

- <a href="#VUID-vkEndCommandBuffer-commandBuffer-00062"
  id="VUID-vkEndCommandBuffer-commandBuffer-00062"></a>
  VUID-vkEndCommandBuffer-commandBuffer-00062  
  If `commandBuffer` is a secondary command buffer, there **must** not
  be an outstanding
  [vkCmdDebugMarkerBeginEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDebugMarkerBeginEXT.html) command
  recorded to `commandBuffer` that has not previously been ended by a
  call to [vkCmdDebugMarkerEndEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDebugMarkerEndEXT.html)

Valid Usage (Implicit)

- <a href="#VUID-vkEndCommandBuffer-commandBuffer-parameter"
  id="VUID-vkEndCommandBuffer-commandBuffer-parameter"></a>
  VUID-vkEndCommandBuffer-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized

- Host access to the `VkCommandPool` that `commandBuffer` was allocated
  from **must** be externally synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_INVALID_VIDEO_STD_PARAMETERS_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkEndCommandBuffer"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
