# vkSetDebugUtilsObjectTagEXT(3) Manual Page

## Name

vkSetDebugUtilsObjectTagEXT - Attach arbitrary data to an object



## <a href="#_c_specification" class="anchor"></a>C Specification

``` c
// Provided by VK_EXT_debug_utils
VkResult vkSetDebugUtilsObjectTagEXT(
    VkDevice                                    device,
    const VkDebugUtilsObjectTagInfoEXT*         pTagInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device that created the object.

- `pTagInfo` is a pointer to a
  [VkDebugUtilsObjectTagInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsObjectTagInfoEXT.html)
  structure specifying parameters of the tag to attach to the object.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkSetDebugUtilsObjectTagEXT-pNameInfo-07875"
  id="VUID-vkSetDebugUtilsObjectTagEXT-pNameInfo-07875"></a>
  VUID-vkSetDebugUtilsObjectTagEXT-pNameInfo-07875  
  If `pNameInfo->objectHandle` is the valid handle of an instance-level
  object, the [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) identified by `device` **must**
  be a descendent of the same [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) as the
  object identified by `pNameInfo->objectHandle`

- <a href="#VUID-vkSetDebugUtilsObjectTagEXT-pNameInfo-07876"
  id="VUID-vkSetDebugUtilsObjectTagEXT-pNameInfo-07876"></a>
  VUID-vkSetDebugUtilsObjectTagEXT-pNameInfo-07876  
  If `pNameInfo->objectHandle` is the valid handle of a
  physical-device-level object, the [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) identified
  by `device` **must** be a descendant of the same
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) as the object identified by
  `pNameInfo->objectHandle`

- <a href="#VUID-vkSetDebugUtilsObjectTagEXT-pNameInfo-07877"
  id="VUID-vkSetDebugUtilsObjectTagEXT-pNameInfo-07877"></a>
  VUID-vkSetDebugUtilsObjectTagEXT-pNameInfo-07877  
  If `pNameInfo->objectHandle` is the valid handle of a device-level
  object, that object **must** be a descendent of the
  [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) identified by `device`

Valid Usage (Implicit)

- <a href="#VUID-vkSetDebugUtilsObjectTagEXT-device-parameter"
  id="VUID-vkSetDebugUtilsObjectTagEXT-device-parameter"></a>
  VUID-vkSetDebugUtilsObjectTagEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkSetDebugUtilsObjectTagEXT-pTagInfo-parameter"
  id="VUID-vkSetDebugUtilsObjectTagEXT-pTagInfo-parameter"></a>
  VUID-vkSetDebugUtilsObjectTagEXT-pTagInfo-parameter  
  `pTagInfo` **must** be a valid pointer to a valid
  [VkDebugUtilsObjectTagInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsObjectTagInfoEXT.html)
  structure

Host Synchronization

- Host access to `pTagInfo->objectHandle` **must** be externally
  synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_debug_utils](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_utils.html),
[VkDebugUtilsObjectTagInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsObjectTagInfoEXT.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkSetDebugUtilsObjectTagEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
