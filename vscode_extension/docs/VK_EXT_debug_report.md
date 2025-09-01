# VK\_EXT\_debug\_report(3) Manual Page

## Name

VK\_EXT\_debug\_report - instance extension



## [](#_registered_extension_number)Registered Extension Number

12

## [](#_revision)Revision

10

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_1

## [](#_deprecation_state)Deprecation State

- *Deprecated* by [VK\_EXT\_debug\_utils](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_utils.html) extension

## [](#_special_use)Special Use

- [Debugging tools](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Courtney Goeltzenleuchter [\[GitHub\]courtney-g](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_debug_report%5D%20%40courtney-g%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_debug_report%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2020-12-14

**IP Status**

No known IP claims.

**Contributors**

- Courtney Goeltzenleuchter, LunarG
- Dan Ginsburg, Valve
- Jon Ashburn, LunarG
- Mark Lobodzinski, LunarG

## [](#_description)Description

Due to the nature of the Vulkan interface, there is very little error information available to the developer and application. By enabling optional validation layers and using the `VK_EXT_debug_report` extension, developers **can** obtain much more detailed feedback on the application’s use of Vulkan. This extension defines a way for layers and the implementation to call back to the application for events of interest to the application.

## [](#_new_object_types)New Object Types

- [VkDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportCallbackEXT.html)

## [](#_new_commands)New Commands

- [vkCreateDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDebugReportCallbackEXT.html)
- [vkDebugReportMessageEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDebugReportMessageEXT.html)
- [vkDestroyDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyDebugReportCallbackEXT.html)

## [](#_new_structures)New Structures

- Extending [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstanceCreateInfo.html):
  
  - [VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportCallbackCreateInfoEXT.html)

## [](#_new_function_pointers)New Function Pointers

- [PFN\_vkDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/PFN_vkDebugReportCallbackEXT.html)

## [](#_new_enums)New Enums

- [VkDebugReportFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportFlagBitsEXT.html)
- [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportObjectTypeEXT.html)

## [](#_new_bitmasks)New Bitmasks

- [VkDebugReportFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportFlagsEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_DEBUG_REPORT_EXTENSION_NAME`
- `VK_EXT_DEBUG_REPORT_SPEC_VERSION`
- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html):
  
  - `VK_OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT`
- Extending [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html):
  
  - `VK_ERROR_VALIDATION_FAILED_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DEBUG_REPORT_CALLBACK_CREATE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_DEBUG_REPORT_CREATE_INFO_EXT`

If [Vulkan Version 1.1](#versions-1.1) is supported:

- Extending [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportObjectTypeEXT.html):
  
  - `VK_DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_EXT`
  - `VK_DEBUG_REPORT_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_EXT`

## [](#_examples)Examples

`VK_EXT_debug_report` allows an application to register multiple callbacks with the validation layers. Some callbacks may log the information to a file, others may cause a debug break point or other application-defined behavior. An application **can** register callbacks even when no validation layers are enabled, but they will only be called for loader and, if implemented, driver events.

To capture events that occur while creating or destroying an instance an application **can** link a [VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportCallbackCreateInfoEXT.html) structure to the `pNext` chain of the [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstanceCreateInfo.html) structure passed to [vkCreateInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateInstance.html).

Example uses: Create three callback objects. One will log errors and warnings to the debug console using Windows `OutputDebugString`. The second will cause the debugger to break at that callback when an error happens and the third will log warnings to stdout.

```c++
    VkResult res;
    VkDebugReportCallbackEXT cb1, cb2, cb3;

    VkDebugReportCallbackCreateInfoEXT callback1 = {
        .sType = VK_STRUCTURE_TYPE_DEBUG_REPORT_CALLBACK_CREATE_INFO_EXT,
        .pNext = NULL,
        .flags = VK_DEBUG_REPORT_ERROR_BIT_EXT |
                 VK_DEBUG_REPORT_WARNING_BIT_EXT,
        .pfnCallback = myOutputDebugString,
        .pUserData = NULL
    };
    res = vkCreateDebugReportCallbackEXT(instance, &callback1, &cb1);
    if (res != VK_SUCCESS)
       /* Do error handling for VK_ERROR_OUT_OF_MEMORY */

    callback.flags = VK_DEBUG_REPORT_ERROR_BIT_EXT;
    callback.pfnCallback = myDebugBreak;
    callback.pUserData = NULL;
    res = vkCreateDebugReportCallbackEXT(instance, &callback, &cb2);
    if (res != VK_SUCCESS)
       /* Do error handling for VK_ERROR_OUT_OF_MEMORY */

    VkDebugReportCallbackCreateInfoEXT callback3 = {
        .sType = VK_STRUCTURE_TYPE_DEBUG_REPORT_CALLBACK_CREATE_INFO_EXT,
        .pNext = NULL,
        .flags = VK_DEBUG_REPORT_WARNING_BIT_EXT,
        .pfnCallback = mystdOutLogger,
        .pUserData = NULL
    };
    res = vkCreateDebugReportCallbackEXT(instance, &callback3, &cb3);
    if (res != VK_SUCCESS)
       /* Do error handling for VK_ERROR_OUT_OF_MEMORY */

    ...

    /* remove callbacks when cleaning up */
    vkDestroyDebugReportCallbackEXT(instance, cb1);
    vkDestroyDebugReportCallbackEXT(instance, cb2);
    vkDestroyDebugReportCallbackEXT(instance, cb3);
```

Note

In the initial release of the `VK_EXT_debug_report` extension, the token `VK_STRUCTURE_TYPE_DEBUG_REPORT_CREATE_INFO_EXT` was used. Starting in version 2 of the extension branch, `VK_STRUCTURE_TYPE_DEBUG_REPORT_CALLBACK_CREATE_INFO_EXT` is used instead for consistency with Vulkan naming rules. The older enum is still available for backwards compatibility.

Note

In the initial release of the `VK_EXT_debug_report` extension, the token `VK_DEBUG_REPORT_OBJECT_TYPE_DEBUG_REPORT_EXT` was used. Starting in version 8 of the extension branch, `VK_DEBUG_REPORT_OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT_EXT` is used instead for consistency with Vulkan naming rules. The older enum is still available for backwards compatibility.

## [](#_issues)Issues

1\) What is the hierarchy / seriousness of the message flags? E.g. `ERROR` &gt; `WARN` &gt; `PERF_WARN` …​

**RESOLVED**: There is no specific hierarchy. Each bit is independent and should be checked via bitwise AND. For example:

```c++
    if (localFlags & VK_DEBUG_REPORT_ERROR_BIT_EXT) {
        process error message
    }
    if (localFlags & VK_DEBUG_REPORT_DEBUG_BIT_EXT) {
        process debug message
    }
```

The validation layers do use them in a hierarchical way (`ERROR` &gt; `WARN` &gt; `PERF`, `WARN` &gt; `DEBUG` &gt; `INFO`) and they (at least at the time of this writing) only set one bit at a time. But it is not a requirement of this extension.

It is possible that a layer may intercept and change, or augment the flags with extension values the application’s debug report handler may not be familiar with, so it is important to treat each flag independently.

2\) Should there be a VU requiring [VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportCallbackCreateInfoEXT.html)::`flags` to be non-zero?

**RESOLVED**: It may not be very useful, but we do not need VU statement requiring the `VkDebugReportCallbackCreateInfoEXT`::`msgFlags` at create-time to be non-zero. One can imagine that apps may prefer it as it allows them to set the mask as desired - including nothing - at runtime without having to check.

3\) What is the difference between `VK_DEBUG_REPORT_DEBUG_BIT_EXT` and `VK_DEBUG_REPORT_INFORMATION_BIT_EXT`?

**RESOLVED**: `VK_DEBUG_REPORT_DEBUG_BIT_EXT` specifies information that could be useful debugging the Vulkan implementation itself.

4\) How do you compare handles returned by the debug\_report callback to the application’s handles?

**RESOLVED**: Due to the different nature of dispatchable and nondispatchable handles there is no generic way (that we know of) that works for common compilers with 32bit, 64bit, C and C++. We recommend applications use the same cast that the validation layers use:

\+

```c++
reinterpret_cast<uint64_t &>(dispatchableHandle)
(uint64_t)(nondispatchableHandle)
```

\+ This does require that the application treat dispatchable and nondispatchable handles differently.

## [](#_version_history)Version History

- Revision 1, 2015-05-20 (Courtney Goetzenleuchter)
  
  - Initial draft, based on LunarG KHR spec, other KHR specs
- Revision 2, 2016-02-16 (Courtney Goetzenleuchter)
  
  - Update usage, documentation
- Revision 3, 2016-06-14 (Courtney Goetzenleuchter)
  
  - Update VK\_EXT\_DEBUG\_REPORT\_SPEC\_VERSION to indicate added support for vkCreateInstance and vkDestroyInstance
- Revision 4, 2016-12-08 (Mark Lobodzinski)
  
  - Added Display\_KHR, DisplayModeKHR extension objects
  - Added ObjectTable\_NVX, IndirectCommandsLayout\_NVX extension objects
  - Bumped spec revision
  - Retroactively added version history
- Revision 5, 2017-01-31 (Baldur Karlsson)
  
  - Moved definition of [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportObjectTypeEXT.html) from debug marker chapter
- Revision 6, 2017-01-31 (Baldur Karlsson)
  
  - Added VK\_DEBUG\_REPORT\_OBJECT\_TYPE\_DESCRIPTOR\_UPDATE\_TEMPLATE\_KHR\_EXT
- Revision 7, 2017-04-20 (Courtney Goeltzenleuchter)
  
  - Clarify wording and address questions from developers.
- Revision 8, 2017-04-21 (Courtney Goeltzenleuchter)
  
  - Remove unused enum VkDebugReportErrorEXT
- Revision 9, 2017-09-12 (Tobias Hector)
  
  - Added interactions with Vulkan 1.1
- Revision 10, 2020-12-14 (Courtney Goetzenleuchter)
  
  - Add issue 4 discussing matching handles returned by the extension, based on suggestion in public issue 368.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_debug_report).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0