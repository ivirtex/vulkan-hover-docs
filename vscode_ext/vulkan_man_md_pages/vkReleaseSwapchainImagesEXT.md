# vkReleaseSwapchainImagesEXT(3) Manual Page

## Name

vkReleaseSwapchainImagesEXT - Release previously acquired but unused
images



## <a href="#_c_specification" class="anchor"></a>C Specification

To release images previously acquired through
[vkAcquireNextImage2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireNextImage2KHR.html) or
[vkAcquireNextImageKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireNextImageKHR.html), call:

``` c
// Provided by VK_EXT_swapchain_maintenance1
VkResult vkReleaseSwapchainImagesEXT(
    VkDevice                                    device,
    const VkReleaseSwapchainImagesInfoEXT*      pReleaseInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device associated with
  [VkReleaseSwapchainImagesInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkReleaseSwapchainImagesInfoEXT.html)::`swapchain`.

- `pReleaseInfo` is a pointer to a
  [VkReleaseSwapchainImagesInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkReleaseSwapchainImagesInfoEXT.html)
  structure containing parameters of the release.

## <a href="#_description" class="anchor"></a>Description

Only images that are not in use by the device **can** be released.

Releasing images is a read-only operation that will not affect the
content of the released images. Upon reacquiring the image, the image
contents and its layout will be the same as they were prior to releasing
it. However, if a mechanism other than Vulkan is used to modify the
platform window associated with the swapchain, the content of all
presentable images in the swapchain becomes undefined.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>This functionality is useful during swapchain recreation, where
acquired images from the old swapchain can be released instead of
presented.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-vkReleaseSwapchainImagesEXT-device-parameter"
  id="VUID-vkReleaseSwapchainImagesEXT-device-parameter"></a>
  VUID-vkReleaseSwapchainImagesEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkReleaseSwapchainImagesEXT-pReleaseInfo-parameter"
  id="VUID-vkReleaseSwapchainImagesEXT-pReleaseInfo-parameter"></a>
  VUID-vkReleaseSwapchainImagesEXT-pReleaseInfo-parameter  
  `pReleaseInfo` **must** be a valid pointer to a valid
  [VkReleaseSwapchainImagesInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkReleaseSwapchainImagesInfoEXT.html)
  structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_SURFACE_LOST_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_swapchain_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_swapchain_maintenance1.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkReleaseSwapchainImagesInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkReleaseSwapchainImagesInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkReleaseSwapchainImagesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
