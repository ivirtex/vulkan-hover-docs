# vkBindImageMemory2(3) Manual Page

## Name

vkBindImageMemory2 - Bind device memory to image objects



## [](#_c_specification)C Specification

To attach memory to image objects for one or more images at a time, call:

```c++
// Provided by VK_VERSION_1_1
VkResult vkBindImageMemory2(
    VkDevice                                    device,
    uint32_t                                    bindInfoCount,
    const VkBindImageMemoryInfo*                pBindInfos);
```

or the equivalent command

```c++
// Provided by VK_KHR_bind_memory2
VkResult vkBindImageMemory2KHR(
    VkDevice                                    device,
    uint32_t                                    bindInfoCount,
    const VkBindImageMemoryInfo*                pBindInfos);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the images and memory.
- `bindInfoCount` is the number of elements in `pBindInfos`.
- `pBindInfos` is a pointer to an array of [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindImageMemoryInfo.html) structures, describing images and memory to bind.

## [](#_description)Description

On some implementations, it **may** be more efficient to batch memory bindings into a single command.

If the [`maintenance6`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance6) feature is enabled, this command **must** attempt to perform all of the memory binding operations described by `pBindInfos`, and **must** not early exit on the first failure.

If any of the memory binding operations described by `pBindInfos` fail, the [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html) returned by this command **must** be the return value of any one of the memory binding operations which did not return `VK_SUCCESS`.

Note

If the `vkBindImageMemory2` command failed, [VkBindMemoryStatus](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindMemoryStatus.html) structures were not included in the `pNext` chains of each element of `pBindInfos`, and `bindInfoCount` was greater than one, then the images referenced by `pBindInfos` will be in an indeterminate state, and must not be used.

Applications should destroy these images.

Valid Usage

- [](#VUID-vkBindImageMemory2-pBindInfos-02858)VUID-vkBindImageMemory2-pBindInfos-02858  
  If any [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindImageMemoryInfo.html)::`image` was created with `VK_IMAGE_CREATE_DISJOINT_BIT` then all planes of [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindImageMemoryInfo.html)::`image` **must** be bound individually in separate `pBindInfos`
- [](#VUID-vkBindImageMemory2-pBindInfos-04006)VUID-vkBindImageMemory2-pBindInfos-04006  
  `pBindInfos` **must** not refer to the same image subresource more than once

Valid Usage (Implicit)

- [](#VUID-vkBindImageMemory2-device-parameter)VUID-vkBindImageMemory2-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkBindImageMemory2-pBindInfos-parameter)VUID-vkBindImageMemory2-pBindInfos-parameter  
  `pBindInfos` **must** be a valid pointer to an array of `bindInfoCount` valid [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindImageMemoryInfo.html) structures
- [](#VUID-vkBindImageMemory2-bindInfoCount-arraylength)VUID-vkBindImageMemory2-bindInfoCount-arraylength  
  `bindInfoCount` **must** be greater than `0`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindImageMemoryInfo.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkBindImageMemory2).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0