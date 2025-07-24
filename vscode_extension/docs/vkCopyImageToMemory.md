# vkCopyImageToMemory(3) Manual Page

## Name

vkCopyImageToMemory - Copy image data into host memory



## [](#_c_specification)C Specification

To copy data from an image object to host memory, call:

```c++
// Provided by VK_VERSION_1_4
VkResult vkCopyImageToMemory(
    VkDevice                                    device,
    const VkCopyImageToMemoryInfo*              pCopyImageToMemoryInfo);
```

or the equivalent command

```c++
// Provided by VK_EXT_host_image_copy
VkResult vkCopyImageToMemoryEXT(
    VkDevice                                    device,
    const VkCopyImageToMemoryInfo*              pCopyImageToMemoryInfo);
```

## [](#_parameters)Parameters

- `device` is the device which owns `pCopyImageToMemoryInfo->srcImage`.
- `pCopyImageToMemoryInfo` is a pointer to a [VkCopyImageToMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToMemoryInfo.html) structure describing the copy parameters.

## [](#_description)Description

This command is functionally similar to [vkCmdCopyImageToBuffer2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyImageToBuffer2.html), except it is executed on the host and writes to host memory instead of a buffer. The memory of `pCopyImageToMemoryInfo->srcImage` is accessed by the host as if [coherent](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-coherent).

Note

If the device has written to the image memory, it is not automatically made available to the host. Before this copy command can be called, a memory barrier for this image **must** have been issued on the device with the second [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) including `VK_PIPELINE_STAGE_HOST_BIT` and `VK_ACCESS_HOST_READ_BIT`.

Valid Usage (Implicit)

- [](#VUID-vkCopyImageToMemory-device-parameter)VUID-vkCopyImageToMemory-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCopyImageToMemory-pCopyImageToMemoryInfo-parameter)VUID-vkCopyImageToMemory-pCopyImageToMemoryInfo-parameter  
  `pCopyImageToMemoryInfo` **must** be a valid pointer to a valid [VkCopyImageToMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToMemoryInfo.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INITIALIZATION_FAILED`
- `VK_ERROR_MEMORY_MAP_FAILED`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_host\_image\_copy](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_host_image_copy.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkCopyImageToMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToMemoryInfo.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCopyImageToMemory)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0