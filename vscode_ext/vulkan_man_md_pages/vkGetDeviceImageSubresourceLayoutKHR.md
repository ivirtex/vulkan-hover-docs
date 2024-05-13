# vkGetDeviceImageSubresourceLayoutKHR(3) Manual Page

## Name

vkGetDeviceImageSubresourceLayoutKHR - Retrieve information about an
image subresource without an image object



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the memory layout of an image subresource, without an image
object, call:

``` c
// Provided by VK_KHR_maintenance5
void vkGetDeviceImageSubresourceLayoutKHR(
    VkDevice                                    device,
    const VkDeviceImageSubresourceInfoKHR*      pInfo,
    VkSubresourceLayout2KHR*                    pLayout);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the image.

- `pInfo` is a pointer to a
  [VkDeviceImageSubresourceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceImageSubresourceInfoKHR.html)
  structure containing parameters required for the subresource layout
  query.

- `pLayout` is a pointer to a
  [VkSubresourceLayout2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout2KHR.html) structure in
  which the layout is returned.

## <a href="#_description" class="anchor"></a>Description

`vkGetDeviceImageSubresourceLayoutKHR` behaves similarly to
[vkGetImageSubresourceLayout2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSubresourceLayout2KHR.html),
but uses a [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) structure to
specify the image rather than a [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) object.

Valid Usage (Implicit)

- <a href="#VUID-vkGetDeviceImageSubresourceLayoutKHR-device-parameter"
  id="VUID-vkGetDeviceImageSubresourceLayoutKHR-device-parameter"></a>
  VUID-vkGetDeviceImageSubresourceLayoutKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetDeviceImageSubresourceLayoutKHR-pInfo-parameter"
  id="VUID-vkGetDeviceImageSubresourceLayoutKHR-pInfo-parameter"></a>
  VUID-vkGetDeviceImageSubresourceLayoutKHR-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkDeviceImageSubresourceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceImageSubresourceInfoKHR.html)
  structure

- <a href="#VUID-vkGetDeviceImageSubresourceLayoutKHR-pLayout-parameter"
  id="VUID-vkGetDeviceImageSubresourceLayoutKHR-pLayout-parameter"></a>
  VUID-vkGetDeviceImageSubresourceLayoutKHR-pLayout-parameter  
  `pLayout` **must** be a valid pointer to a
  [VkSubresourceLayout2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout2KHR.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance5](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance5.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkDeviceImageSubresourceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceImageSubresourceInfoKHR.html),
[VkSubresourceLayout2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDeviceImageSubresourceLayoutKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
