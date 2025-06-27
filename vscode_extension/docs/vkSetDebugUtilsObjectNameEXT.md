# vkSetDebugUtilsObjectNameEXT(3) Manual Page

## Name

vkSetDebugUtilsObjectNameEXT - Give an application-defined name to an
object



## <a href="#_c_specification" class="anchor"></a>C Specification

An object can be given an application-defined name by calling:

``` c
// Provided by VK_EXT_debug_utils
VkResult vkSetDebugUtilsObjectNameEXT(
    VkDevice                                    device,
    const VkDebugUtilsObjectNameInfoEXT*        pNameInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device that is associated with the named object passed
  in via `objectHandle`.

- `pNameInfo` is a pointer to a
  [VkDebugUtilsObjectNameInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsObjectNameInfoEXT.html)
  structure specifying parameters of the name to set on the object.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-02587"
  id="VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-02587"></a>
  VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-02587  
  `pNameInfo->objectType` **must** not be `VK_OBJECT_TYPE_UNKNOWN`

- <a href="#VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-02588"
  id="VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-02588"></a>
  VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-02588  
  `pNameInfo->objectHandle` **must** not be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-07872"
  id="VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-07872"></a>
  VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-07872  
  If `pNameInfo->objectHandle` is the valid handle of an instance-level
  object, the [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) identified by `device` **must**
  be a descendent of the same [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) as the
  object identified by `pNameInfo->objectHandle`

- <a href="#VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-07873"
  id="VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-07873"></a>
  VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-07873  
  If `pNameInfo->objectHandle` is the valid handle of a
  physical-device-level object, the [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) identified
  by `device` **must** be a descendant of the same
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) as the object identified by
  `pNameInfo->objectHandle`

- <a href="#VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-07874"
  id="VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-07874"></a>
  VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-07874  
  If `pNameInfo->objectHandle` is the valid handle of a device-level
  object, that object **must** be a descendent of the
  [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) identified by `device`

Valid Usage (Implicit)

- <a href="#VUID-vkSetDebugUtilsObjectNameEXT-device-parameter"
  id="VUID-vkSetDebugUtilsObjectNameEXT-device-parameter"></a>
  VUID-vkSetDebugUtilsObjectNameEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-parameter"
  id="VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-parameter"></a>
  VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-parameter  
  `pNameInfo` **must** be a valid pointer to a valid
  [VkDebugUtilsObjectNameInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsObjectNameInfoEXT.html)
  structure

Host Synchronization

- Host access to `pNameInfo->objectHandle` **must** be externally
  synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_debug_utils](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_utils.html),
[VkDebugUtilsObjectNameInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsObjectNameInfoEXT.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkSetDebugUtilsObjectNameEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
