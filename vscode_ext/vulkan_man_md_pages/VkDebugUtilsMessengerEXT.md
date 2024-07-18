# VkDebugUtilsMessengerEXT(3) Manual Page

## Name

VkDebugUtilsMessengerEXT - Opaque handle to a debug messenger object



## <a href="#_c_specification" class="anchor"></a>C Specification

A `VkDebugUtilsMessengerEXT` is a messenger object which handles passing
along debug messages to a provided debug callback.

``` c
// Provided by VK_EXT_debug_utils
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkDebugUtilsMessengerEXT)
```

## <a href="#_description" class="anchor"></a>Description

The debug messenger will provide detailed feedback on the application’s
use of Vulkan when events of interest occur. When an event of interest
does occur, the debug messenger will submit a debug message to the debug
callback that was provided during its creation. Additionally, the debug
messenger is responsible with filtering out debug messages that the
callback is not interested in and will only provide desired debug
messages.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_DEFINE_NON_DISPATCHABLE_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html),
[VK_EXT_debug_utils](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_utils.html),
[vkCreateDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDebugUtilsMessengerEXT.html),
[vkDestroyDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyDebugUtilsMessengerEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDebugUtilsMessengerEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
