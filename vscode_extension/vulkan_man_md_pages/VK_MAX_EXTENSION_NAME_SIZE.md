# VK_MAX_EXTENSION_NAME_SIZE(3) Manual Page

## Name

VK_MAX_EXTENSION_NAME_SIZE - Maximum length of a layer of extension name
string



## <a href="#_c_specification" class="anchor"></a>C Specification

`VK_MAX_EXTENSION_NAME_SIZE` is the length in `char` values of an array
containing a layer or extension name string, as returned in
[VkLayerProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerProperties.html)::`layerName`,
[VkExtensionProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtensionProperties.html)::`extensionName`,
and other queries.

``` c
#define VK_MAX_EXTENSION_NAME_SIZE        256U
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_MAX_EXTENSION_NAME_SIZE"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
