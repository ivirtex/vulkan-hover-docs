# VK_NULL_HANDLE(3) Manual Page

## Name

VK_NULL_HANDLE - Reserved non-valid object handle



## <a href="#_c_specification" class="anchor"></a>C Specification

`VK_NULL_HANDLE` is a reserved value representing a non-valid object
handle. It may be passed to and returned from Vulkan commands only when
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fundamentals-validusage-handles"
target="_blank" rel="noopener">specifically allowed</a>.

``` c
// Provided by VK_VERSION_1_0

#ifndef VK_DEFINE_NON_DISPATCHABLE_HANDLE
    #if (VK_USE_64_BIT_PTR_DEFINES==1)
        #if (defined(__cplusplus) && (__cplusplus >= 201103L)) || (defined(_MSVC_LANG) && (_MSVC_LANG >= 201103L))
            #define VK_NULL_HANDLE nullptr
        #else
            #define VK_NULL_HANDLE ((void*)0)
        #endif
    #else
        #define VK_NULL_HANDLE 0ULL
    #endif
#endif
#ifndef VK_NULL_HANDLE
    #define VK_NULL_HANDLE 0
#endif
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_DEFINE_NON_DISPATCHABLE_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html),
[VK_USE_64_BIT_PTR_DEFINES](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_USE_64_BIT_PTR_DEFINES.html),
[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NULL_HANDLE"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
