# vkBindBufferMemory2(3) Manual Page

## Name

vkBindBufferMemory2 - Bind device memory to buffer objects



## [](#_c_specification)C Specification

To attach memory to buffer objects for one or more buffers at a time, call:

```c++
// Provided by VK_VERSION_1_1
VkResult vkBindBufferMemory2(
    VkDevice                                    device,
    uint32_t                                    bindInfoCount,
    const VkBindBufferMemoryInfo*               pBindInfos);
```

or the equivalent command

```c++
// Provided by VK_KHR_bind_memory2
VkResult vkBindBufferMemory2KHR(
    VkDevice                                    device,
    uint32_t                                    bindInfoCount,
    const VkBindBufferMemoryInfo*               pBindInfos);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the buffers and memory.
- `bindInfoCount` is the number of elements in `pBindInfos`.
- `pBindInfos` is a pointer to an array of `bindInfoCount` [VkBindBufferMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindBufferMemoryInfo.html) structures describing buffers and memory to bind.

## [](#_description)Description

On some implementations, it **may** be more efficient to batch memory bindings into a single command.

If the [`maintenance6`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance6) feature is enabled, this command **must** attempt to perform all of the memory binding operations described by `pBindInfos`, and **must** not early exit on the first failure.

If any of the memory binding operations described by `pBindInfos` fail, the [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html) returned by this command **must** be the return value of any one of the memory binding operations which did not return `VK_SUCCESS`.

Note

If the `vkBindBufferMemory2` command failed, [VkBindMemoryStatus](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindMemoryStatus.html) structures were not included in the `pNext` chains of each element of `pBindInfos`, and `bindInfoCount` was greater than one, then the buffers referenced by `pBindInfos` will be in an indeterminate state, and must not be used.

Applications should destroy these buffers.

Valid Usage (Implicit)

- [](#VUID-vkBindBufferMemory2-device-parameter)VUID-vkBindBufferMemory2-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkBindBufferMemory2-pBindInfos-parameter)VUID-vkBindBufferMemory2-pBindInfos-parameter  
  `pBindInfos` **must** be a valid pointer to an array of `bindInfoCount` valid [VkBindBufferMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindBufferMemoryInfo.html) structures
- [](#VUID-vkBindBufferMemory2-bindInfoCount-arraylength)VUID-vkBindBufferMemory2-bindInfoCount-arraylength  
  `bindInfoCount` **must** be greater than `0`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR`

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkBindBufferMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindBufferMemoryInfo.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkBindBufferMemory2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0