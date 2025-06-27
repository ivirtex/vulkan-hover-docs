# vkGetSwapchainImagesKHR(3) Manual Page

## Name

vkGetSwapchainImagesKHR - Obtain the array of presentable images
associated with a swapchain



## <a href="#_c_specification" class="anchor"></a>C Specification

To obtain the array of presentable images associated with a swapchain,
call:

``` c
// Provided by VK_KHR_swapchain
VkResult vkGetSwapchainImagesKHR(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain,
    uint32_t*                                   pSwapchainImageCount,
    VkImage*                                    pSwapchainImages);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device associated with `swapchain`.

- `swapchain` is the swapchain to query.

- `pSwapchainImageCount` is a pointer to an integer related to the
  number of presentable images available or queried, as described below.

- `pSwapchainImages` is either `NULL` or a pointer to an array of
  `VkImage` handles.

## <a href="#_description" class="anchor"></a>Description

If `pSwapchainImages` is `NULL`, then the number of presentable images
for `swapchain` is returned in `pSwapchainImageCount`. Otherwise,
`pSwapchainImageCount` **must** point to a variable set by the
application to the number of elements in the `pSwapchainImages` array,
and on return the variable is overwritten with the number of structures
actually written to `pSwapchainImages`. If the value of
`pSwapchainImageCount` is less than the number of presentable images for
`swapchain`, at most `pSwapchainImageCount` structures will be written,
and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to
indicate that not all the available presentable images were returned.

Valid Usage (Implicit)

- <a href="#VUID-vkGetSwapchainImagesKHR-device-parameter"
  id="VUID-vkGetSwapchainImagesKHR-device-parameter"></a>
  VUID-vkGetSwapchainImagesKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetSwapchainImagesKHR-swapchain-parameter"
  id="VUID-vkGetSwapchainImagesKHR-swapchain-parameter"></a>
  VUID-vkGetSwapchainImagesKHR-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)
  handle

- <a href="#VUID-vkGetSwapchainImagesKHR-pSwapchainImageCount-parameter"
  id="VUID-vkGetSwapchainImagesKHR-pSwapchainImageCount-parameter"></a>
  VUID-vkGetSwapchainImagesKHR-pSwapchainImageCount-parameter  
  `pSwapchainImageCount` **must** be a valid pointer to a `uint32_t`
  value

- <a href="#VUID-vkGetSwapchainImagesKHR-pSwapchainImages-parameter"
  id="VUID-vkGetSwapchainImagesKHR-pSwapchainImages-parameter"></a>
  VUID-vkGetSwapchainImagesKHR-pSwapchainImages-parameter  
  If the value referenced by `pSwapchainImageCount` is not `0`, and
  `pSwapchainImages` is not `NULL`, `pSwapchainImages` **must** be a
  valid pointer to an array of `pSwapchainImageCount`
  [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handles

- <a href="#VUID-vkGetSwapchainImagesKHR-swapchain-parent"
  id="VUID-vkGetSwapchainImagesKHR-swapchain-parent"></a>
  VUID-vkGetSwapchainImagesKHR-swapchain-parent  
  `swapchain` **must** have been created, allocated, or retrieved from
  `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetSwapchainImagesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
