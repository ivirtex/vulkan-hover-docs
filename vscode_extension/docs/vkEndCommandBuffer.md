# vkEndCommandBuffer(3) Manual Page

## Name

vkEndCommandBuffer - Finish recording a command buffer



## [](#_c_specification)C Specification

To complete recording of a command buffer, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkEndCommandBuffer(
    VkCommandBuffer                             commandBuffer);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer to complete recording.

## [](#_description)Description

The command buffer **must** have been in the [recording state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle), and, if successful, is moved to the [executable state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle).

If there was an error during recording, the application will be notified by an unsuccessful return code returned by `vkEndCommandBuffer`, and the command buffer will be moved to the [invalid state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle).

In case the application recorded one or more [video encode operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-encode-operations) into the command buffer, implementations **may** return the `VK_ERROR_INVALID_VIDEO_STD_PARAMETERS_KHR` error if any of the specified Video Std parameters do not adhere to the syntactic or semantic requirements of the used video compression standard, or if values derived from parameters according to the rules defined by the used video compression standard do not adhere to the capabilities of the video compression standard or the implementation.

Note

Applications **should** not rely on the `VK_ERROR_INVALID_VIDEO_STD_PARAMETERS_KHR` error being returned by any command as a means to verify Video Std parameters, as implementations are not required to report the error in any specific set of cases.

Valid Usage

- [](#VUID-vkEndCommandBuffer-commandBuffer-00059)VUID-vkEndCommandBuffer-commandBuffer-00059  
  `commandBuffer` **must** be in the [recording state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle)
- [](#VUID-vkEndCommandBuffer-commandBuffer-00060)VUID-vkEndCommandBuffer-commandBuffer-00060  
  If `commandBuffer` is a primary command buffer, there **must** not be an active render pass instance
- [](#VUID-vkEndCommandBuffer-commandBuffer-00061)VUID-vkEndCommandBuffer-commandBuffer-00061  
  All queries made [active](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-operation-active) during the recording of `commandBuffer` **must** have been made inactive
- [](#VUID-vkEndCommandBuffer-None-01978)VUID-vkEndCommandBuffer-None-01978  
  Conditional rendering **must** not be [active](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#active-conditional-rendering)
- [](#VUID-vkEndCommandBuffer-None-06991)VUID-vkEndCommandBuffer-None-06991  
  There **must** be no video session object bound
- [](#VUID-vkEndCommandBuffer-commandBuffer-01815)VUID-vkEndCommandBuffer-commandBuffer-01815  
  If `commandBuffer` is a secondary command buffer, there **must** not be an outstanding [vkCmdBeginDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginDebugUtilsLabelEXT.html) command recorded to `commandBuffer` that has not previously been ended by a call to [vkCmdEndDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndDebugUtilsLabelEXT.html)
- [](#VUID-vkEndCommandBuffer-commandBuffer-00062)VUID-vkEndCommandBuffer-commandBuffer-00062  
  If `commandBuffer` is a secondary command buffer, there **must** not be an outstanding [vkCmdDebugMarkerBeginEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDebugMarkerBeginEXT.html) command recorded to `commandBuffer` that has not previously been ended by a call to [vkCmdDebugMarkerEndEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDebugMarkerEndEXT.html)

Valid Usage (Implicit)

- [](#VUID-vkEndCommandBuffer-commandBuffer-parameter)VUID-vkEndCommandBuffer-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INVALID_VIDEO_STD_PARAMETERS_KHR`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkEndCommandBuffer)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0