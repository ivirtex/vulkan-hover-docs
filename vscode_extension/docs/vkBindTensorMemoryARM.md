# vkBindTensorMemoryARM(3) Manual Page

## Name

vkBindTensorMemoryARM - Bind device memory to tensor objects



## [](#_c_specification)C Specification

To attach memory to tensor objects call:

```c++
// Provided by VK_ARM_tensors
VkResult vkBindTensorMemoryARM(
    VkDevice                                    device,
    uint32_t                                    bindInfoCount,
    const VkBindTensorMemoryInfoARM*            pBindInfos);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the buffers and memory.
- `bindInfoCount` is the number of elements in `pBindInfos`.
- `pBindInfos` is a pointer to an array of structures of type [VkBindTensorMemoryInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindTensorMemoryInfoARM.html), describing tensors and memory to bind.

## [](#_description)Description

On some implementations, it **may** be more efficient to batch memory bindings into a single command.

Valid Usage (Implicit)

- [](#VUID-vkBindTensorMemoryARM-device-parameter)VUID-vkBindTensorMemoryARM-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkBindTensorMemoryARM-pBindInfos-parameter)VUID-vkBindTensorMemoryARM-pBindInfos-parameter  
  `pBindInfos` **must** be a valid pointer to an array of `bindInfoCount` valid [VkBindTensorMemoryInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindTensorMemoryInfoARM.html) structures
- [](#VUID-vkBindTensorMemoryARM-bindInfoCount-arraylength)VUID-vkBindTensorMemoryARM-bindInfoCount-arraylength  
  `bindInfoCount` **must** be greater than `0`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkBindTensorMemoryInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindTensorMemoryInfoARM.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkBindTensorMemoryARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0