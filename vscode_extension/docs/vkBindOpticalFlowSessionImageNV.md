# vkBindOpticalFlowSessionImageNV(3) Manual Page

## Name

vkBindOpticalFlowSessionImageNV - Bind image to an optical flow session



## [](#_c_specification)C Specification

To bind a Vulkan image to an optical flow session object, call:

```c++
// Provided by VK_NV_optical_flow
VkResult vkBindOpticalFlowSessionImageNV(
    VkDevice                                    device,
    VkOpticalFlowSessionNV                      session,
    VkOpticalFlowSessionBindingPointNV          bindingPoint,
    VkImageView                                 view,
    VkImageLayout                               layout);
```

## [](#_parameters)Parameters

- `device` is the device which owns the optical flow session object `session`.
- `session` is the optical flow session object to which the image view is to be bound.
- `bindingPoint` specifies the binding point [VkOpticalFlowSessionBindingPointNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionBindingPointNV.html) to which the image view is bound.
- `view` is a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) to be bound.
- `layout` **must** specify the layout that the image subresources accessible from `view` will be in at the time the optical flow vectors are calculated with [vkCmdOpticalFlowExecuteNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdOpticalFlowExecuteNV.html) on a `VkDevice`.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkBindOpticalFlowSessionImageNV-device-parameter)VUID-vkBindOpticalFlowSessionImageNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkBindOpticalFlowSessionImageNV-session-parameter)VUID-vkBindOpticalFlowSessionImageNV-session-parameter  
  `session` **must** be a valid [VkOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionNV.html) handle
- [](#VUID-vkBindOpticalFlowSessionImageNV-bindingPoint-parameter)VUID-vkBindOpticalFlowSessionImageNV-bindingPoint-parameter  
  `bindingPoint` **must** be a valid [VkOpticalFlowSessionBindingPointNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionBindingPointNV.html) value
- [](#VUID-vkBindOpticalFlowSessionImageNV-view-parameter)VUID-vkBindOpticalFlowSessionImageNV-view-parameter  
  If `view` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `view` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) handle
- [](#VUID-vkBindOpticalFlowSessionImageNV-layout-parameter)VUID-vkBindOpticalFlowSessionImageNV-layout-parameter  
  `layout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-vkBindOpticalFlowSessionImageNV-session-parent)VUID-vkBindOpticalFlowSessionImageNV-session-parent  
  `session` **must** have been created, allocated, or retrieved from `device`
- [](#VUID-vkBindOpticalFlowSessionImageNV-view-parent)VUID-vkBindOpticalFlowSessionImageNV-view-parent  
  If `view` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INITIALIZATION_FAILED`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_NV\_optical\_flow](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_optical_flow.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html), [VkOpticalFlowSessionBindingPointNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionBindingPointNV.html), [VkOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkBindOpticalFlowSessionImageNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0