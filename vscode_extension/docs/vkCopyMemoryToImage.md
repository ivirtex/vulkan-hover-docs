# vkCopyMemoryToImage(3) Manual Page

## Name

vkCopyMemoryToImage - Copy data from host memory into an image



## [](#_c_specification)C Specification

To copy data from host memory to an image object, call:

```c++
// Provided by VK_VERSION_1_4
VkResult vkCopyMemoryToImage(
    VkDevice                                    device,
    const VkCopyMemoryToImageInfo*              pCopyMemoryToImageInfo);
```

or the equivalent command

```c++
// Provided by VK_EXT_host_image_copy
VkResult vkCopyMemoryToImageEXT(
    VkDevice                                    device,
    const VkCopyMemoryToImageInfo*              pCopyMemoryToImageInfo);
```

## [](#_parameters)Parameters

- `device` is the device which owns `pCopyMemoryToImageInfo->dstImage`.
- `pCopyMemoryToImageInfo` is a pointer to a [VkCopyMemoryToImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToImageInfo.html) structure describing the copy parameters.

## [](#_description)Description

This command is functionally similar to [vkCmdCopyBufferToImage2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyBufferToImage2.html), except it is executed on the host and reads from host memory instead of a buffer. The memory of `pCopyMemoryToImageInfo->dstImage` is accessed by the host as if [coherent](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-coherent).

Note

Because queue submissions [automatically make host memory visible to the device](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-host-writes), there would not be a need for a memory barrier before using the results of this copy operation on the device.

Valid Usage (Implicit)

- [](#VUID-vkCopyMemoryToImage-device-parameter)VUID-vkCopyMemoryToImage-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCopyMemoryToImage-pCopyMemoryToImageInfo-parameter)VUID-vkCopyMemoryToImage-pCopyMemoryToImageInfo-parameter  
  `pCopyMemoryToImageInfo` **must** be a valid pointer to a valid [VkCopyMemoryToImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToImageInfo.html) structure

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

[VK\_EXT\_host\_image\_copy](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_host_image_copy.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkCopyMemoryToImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToImageInfo.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCopyMemoryToImage).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0