# vkTransitionImageLayout(3) Manual Page

## Name

vkTransitionImageLayout - Perform an image layout transition on the host



## [](#_c_specification)C Specification

To perform an image layout transition on the host, call:

```c++
// Provided by VK_VERSION_1_4
VkResult vkTransitionImageLayout(
    VkDevice                                    device,
    uint32_t                                    transitionCount,
    const VkHostImageLayoutTransitionInfo*      pTransitions);
```

or the equivalent command

```c++
// Provided by VK_EXT_host_image_copy
VkResult vkTransitionImageLayoutEXT(
    VkDevice                                    device,
    uint32_t                                    transitionCount,
    const VkHostImageLayoutTransitionInfo*      pTransitions);
```

## [](#_parameters)Parameters

- `device` is the device which owns `pTransitions`\[i].`image`.
- `transitionCount` is the number of image layout transitions to perform.
- `pTransitions` is a pointer to an array of [VkHostImageLayoutTransitionInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHostImageLayoutTransitionInfo.html) structures specifying the image and [subresource ranges](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-views) within them to transition.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkTransitionImageLayout-device-parameter)VUID-vkTransitionImageLayout-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkTransitionImageLayout-pTransitions-parameter)VUID-vkTransitionImageLayout-pTransitions-parameter  
  `pTransitions` **must** be a valid pointer to an array of `transitionCount` valid [VkHostImageLayoutTransitionInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHostImageLayoutTransitionInfo.html) structures
- [](#VUID-vkTransitionImageLayout-transitionCount-arraylength)VUID-vkTransitionImageLayout-transitionCount-arraylength  
  `transitionCount` **must** be greater than `0`

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

[VK\_EXT\_host\_image\_copy](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_host_image_copy.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkHostImageLayoutTransitionInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHostImageLayoutTransitionInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkTransitionImageLayout)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0