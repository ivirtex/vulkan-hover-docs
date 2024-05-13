# VkDebugReportCallbackCreateInfoEXT(3) Manual Page

## Name

VkDebugReportCallbackCreateInfoEXT - Structure specifying parameters of
a newly created debug report callback



## <a href="#_c_specification" class="anchor"></a>C Specification

The definition of
[VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportCallbackCreateInfoEXT.html)
is:

``` c
// Provided by VK_EXT_debug_report
typedef struct VkDebugReportCallbackCreateInfoEXT {
    VkStructureType                 sType;
    const void*                     pNext;
    VkDebugReportFlagsEXT           flags;
    PFN_vkDebugReportCallbackEXT    pfnCallback;
    void*                           pUserData;
} VkDebugReportCallbackCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkDebugReportFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportFlagBitsEXT.html) specifying
  which event(s) will cause this callback to be called.

- `pfnCallback` is the application callback function to call.

- `pUserData` is user data to be passed to the callback.

## <a href="#_description" class="anchor"></a>Description

For each `VkDebugReportCallbackEXT` that is created the
`VkDebugReportCallbackCreateInfoEXT`::`flags` determine when that
`VkDebugReportCallbackCreateInfoEXT`::`pfnCallback` is called. When an
event happens, the implementation will do a bitwise AND of the event’s
[VkDebugReportFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportFlagBitsEXT.html) flags to each
`VkDebugReportCallbackEXT` object’s flags. For each non-zero result the
corresponding callback will be called. The callback will come directly
from the component that detected the event, unless some other layer
intercepts the calls for its own purposes (filter them in a different
way, log to a system error log, etc.).

An application **may** receive multiple callbacks if multiple
`VkDebugReportCallbackEXT` objects were created. A callback will always
be executed in the same thread as the originating Vulkan call.

A callback may be called from multiple threads simultaneously (if the
application is making Vulkan calls from multiple threads).

Valid Usage (Implicit)

- <a href="#VUID-VkDebugReportCallbackCreateInfoEXT-sType-sType"
  id="VUID-VkDebugReportCallbackCreateInfoEXT-sType-sType"></a>
  VUID-VkDebugReportCallbackCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEBUG_REPORT_CALLBACK_CREATE_INFO_EXT`

- <a href="#VUID-VkDebugReportCallbackCreateInfoEXT-flags-parameter"
  id="VUID-VkDebugReportCallbackCreateInfoEXT-flags-parameter"></a>
  VUID-VkDebugReportCallbackCreateInfoEXT-flags-parameter  
  `flags` **must** be a valid combination of
  [VkDebugReportFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportFlagBitsEXT.html) values

- <a href="#VUID-VkDebugReportCallbackCreateInfoEXT-pfnCallback-parameter"
  id="VUID-VkDebugReportCallbackCreateInfoEXT-pfnCallback-parameter"></a>
  VUID-VkDebugReportCallbackCreateInfoEXT-pfnCallback-parameter  
  `pfnCallback` **must** be a valid
  [PFN_vkDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkDebugReportCallbackEXT.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[PFN_vkDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkDebugReportCallbackEXT.html),
[VK_EXT_debug_report](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_report.html),
[VkDebugReportFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportFlagsEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDebugReportCallbackEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDebugReportCallbackCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
