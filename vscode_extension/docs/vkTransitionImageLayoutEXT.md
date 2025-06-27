# vkTransitionImageLayoutEXT(3) Manual Page

## Name

vkTransitionImageLayoutEXT - Perform an image layout transition on the
host



## <a href="#_c_specification" class="anchor"></a>C Specification

To perform an image layout transition on the host, call:

``` c
// Provided by VK_EXT_host_image_copy
VkResult vkTransitionImageLayoutEXT(
    VkDevice                                    device,
    uint32_t                                    transitionCount,
    const VkHostImageLayoutTransitionInfoEXT*   pTransitions);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device which owns `pTransitions`\[i\].`image`.

- `transitionCount` is the number of image layout transitions to
  perform.

- `pTransitions` is a pointer to an array of
  [VkHostImageLayoutTransitionInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHostImageLayoutTransitionInfoEXT.html)
  structures specifying the image and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views"
  target="_blank" rel="noopener">subresource ranges</a> within them to
  transition.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkTransitionImageLayoutEXT-device-parameter"
  id="VUID-vkTransitionImageLayoutEXT-device-parameter"></a>
  VUID-vkTransitionImageLayoutEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkTransitionImageLayoutEXT-pTransitions-parameter"
  id="VUID-vkTransitionImageLayoutEXT-pTransitions-parameter"></a>
  VUID-vkTransitionImageLayoutEXT-pTransitions-parameter  
  `pTransitions` **must** be a valid pointer to an array of
  `transitionCount` valid
  [VkHostImageLayoutTransitionInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHostImageLayoutTransitionInfoEXT.html)
  structures

- <a href="#VUID-vkTransitionImageLayoutEXT-transitionCount-arraylength"
  id="VUID-vkTransitionImageLayoutEXT-transitionCount-arraylength"></a>
  VUID-vkTransitionImageLayoutEXT-transitionCount-arraylength  
  `transitionCount` **must** be greater than `0`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_INITIALIZATION_FAILED`

- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_MEMORY_MAP_FAILED`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_host_image_copy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_host_image_copy.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkHostImageLayoutTransitionInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHostImageLayoutTransitionInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkTransitionImageLayoutEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
