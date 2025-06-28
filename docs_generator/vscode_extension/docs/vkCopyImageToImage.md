# vkCopyImageToImage(3) Manual Page

## Name

vkCopyImageToImage - Copy image data using the host



## [](#_c_specification)C Specification

To copy data from an image object to another image object using the host, call:

```c++
// Provided by VK_VERSION_1_4
VkResult vkCopyImageToImage(
    VkDevice                                    device,
    const VkCopyImageToImageInfo*               pCopyImageToImageInfo);
```

or the equivalent command

```c++
// Provided by VK_EXT_host_image_copy
VkResult vkCopyImageToImageEXT(
    VkDevice                                    device,
    const VkCopyImageToImageInfo*               pCopyImageToImageInfo);
```

## [](#_parameters)Parameters

- `device` is the device which owns `pCopyImageToImageInfo->srcImage` and `pCopyImageToImageInfo->dstImage`.
- `pCopyImageToImageInfo` is a pointer to a [VkCopyImageToImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToImageInfo.html) structure describing the copy parameters.

## [](#_description)Description

This command is functionally similar to [vkCmdCopyImage2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyImage2.html), except it is executed on the host. The memory of `pCopyImageToImageInfo->srcImage` and `pCopyImageToImageInfo->dstImage` is accessed by the host as if [coherent](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-coherent).

Note

If the device has written to the memory of `pCopyImageToImageInfo->srcImage`, it is not automatically made available to the host. Before this copy command can be called, a memory barrier for this image **must** have been issued on the device with the second [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) including `VK_PIPELINE_STAGE_HOST_BIT` and `VK_ACCESS_HOST_READ_BIT`.

Because queue submissions [automatically make host memory visible to the device](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-host-writes), there would not be a need for a memory barrier before using the results of this copy operation in `pCopyMemoryToImageInfo->dstImage` on the device.

Valid Usage (Implicit)

- [](#VUID-vkCopyImageToImage-device-parameter)VUID-vkCopyImageToImage-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCopyImageToImage-pCopyImageToImageInfo-parameter)VUID-vkCopyImageToImage-pCopyImageToImageInfo-parameter  
  `pCopyImageToImageInfo` **must** be a valid pointer to a valid [VkCopyImageToImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToImageInfo.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INITIALIZATION_FAILED`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_MEMORY_MAP_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_host\_image\_copy](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_host_image_copy.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkCopyImageToImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToImageInfo.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCopyImageToImage)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0