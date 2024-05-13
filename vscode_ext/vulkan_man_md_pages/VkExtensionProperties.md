# VkExtensionProperties(3) Manual Page

## Name

VkExtensionProperties - Structure specifying an extension properties



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkExtensionProperties` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkExtensionProperties {
    char        extensionName[VK_MAX_EXTENSION_NAME_SIZE];
    uint32_t    specVersion;
} VkExtensionProperties;
```

## <a href="#_members" class="anchor"></a>Members

- `extensionName` is an array of `VK_MAX_EXTENSION_NAME_SIZE` `char`
  containing a null-terminated UTF-8 string which is the name of the
  extension.

- `specVersion` is the version of this extension. It is an integer,
  incremented with backward compatible changes.

## <a href="#_description" class="anchor"></a>Description

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html),
[VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html),
[vkEnumerateDeviceExtensionProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumerateDeviceExtensionProperties.html),
[vkEnumerateInstanceExtensionProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumerateInstanceExtensionProperties.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExtensionProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
