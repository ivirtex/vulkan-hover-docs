# vkBindOpticalFlowSessionImageNV(3) Manual Page

## Name

vkBindOpticalFlowSessionImageNV - Bind image to an optical flow session



## <a href="#_c_specification" class="anchor"></a>C Specification

To bind a vulkan image to an optical flow session object, call:

``` c
// Provided by VK_NV_optical_flow
VkResult vkBindOpticalFlowSessionImageNV(
    VkDevice                                    device,
    VkOpticalFlowSessionNV                      session,
    VkOpticalFlowSessionBindingPointNV          bindingPoint,
    VkImageView                                 view,
    VkImageLayout                               layout);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device which owns the optical flow session object
  `session`.

- `session` is the optical flow session object to which the image view
  is to be bound.

- `bindingPoint` specifies the binding point
  [VkOpticalFlowSessionBindingPointNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionBindingPointNV.html)
  to which the image view is bound.

- `view` is a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) to be bound.

- layout **must** specify the layout that the image subresources
  accessible from `view` will be in at the time the optical flow vectors
  are calculated with
  [vkCmdOpticalFlowExecuteNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdOpticalFlowExecuteNV.html) on a
  `VkDevice`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkBindOpticalFlowSessionImageNV-device-parameter"
  id="VUID-vkBindOpticalFlowSessionImageNV-device-parameter"></a>
  VUID-vkBindOpticalFlowSessionImageNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkBindOpticalFlowSessionImageNV-session-parameter"
  id="VUID-vkBindOpticalFlowSessionImageNV-session-parameter"></a>
  VUID-vkBindOpticalFlowSessionImageNV-session-parameter  
  `session` **must** be a valid
  [VkOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionNV.html) handle

- <a href="#VUID-vkBindOpticalFlowSessionImageNV-bindingPoint-parameter"
  id="VUID-vkBindOpticalFlowSessionImageNV-bindingPoint-parameter"></a>
  VUID-vkBindOpticalFlowSessionImageNV-bindingPoint-parameter  
  `bindingPoint` **must** be a valid
  [VkOpticalFlowSessionBindingPointNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionBindingPointNV.html)
  value

- <a href="#VUID-vkBindOpticalFlowSessionImageNV-view-parameter"
  id="VUID-vkBindOpticalFlowSessionImageNV-view-parameter"></a>
  VUID-vkBindOpticalFlowSessionImageNV-view-parameter  
  If `view` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `view`
  **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) handle

- <a href="#VUID-vkBindOpticalFlowSessionImageNV-layout-parameter"
  id="VUID-vkBindOpticalFlowSessionImageNV-layout-parameter"></a>
  VUID-vkBindOpticalFlowSessionImageNV-layout-parameter  
  `layout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

- <a href="#VUID-vkBindOpticalFlowSessionImageNV-session-parent"
  id="VUID-vkBindOpticalFlowSessionImageNV-session-parent"></a>
  VUID-vkBindOpticalFlowSessionImageNV-session-parent  
  `session` **must** have been created, allocated, or retrieved from
  `device`

- <a href="#VUID-vkBindOpticalFlowSessionImageNV-view-parent"
  id="VUID-vkBindOpticalFlowSessionImageNV-view-parent"></a>
  VUID-vkBindOpticalFlowSessionImageNV-view-parent  
  If `view` is a valid handle, it **must** have been created, allocated,
  or retrieved from `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_INITIALIZATION_FAILED`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_optical_flow](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_optical_flow.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html),
[VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html),
[VkOpticalFlowSessionBindingPointNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionBindingPointNV.html),
[VkOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkBindOpticalFlowSessionImageNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
