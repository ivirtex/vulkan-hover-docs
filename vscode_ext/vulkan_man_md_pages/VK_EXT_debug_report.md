# VK_EXT_debug_report(3) Manual Page

## Name

VK_EXT_debug_report - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

12

## <a href="#_revision" class="anchor"></a>Revision

10

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_VERSION_1_1

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Deprecated* by [VK_EXT_debug_utils](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_utils.html)
  extension

## <a href="#_special_use" class="anchor"></a>Special Use

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">Debugging tools</a>

## <a href="#_contact" class="anchor"></a>Contact

- Courtney Goeltzenleuchter <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_debug_report%5D%20@courtney-g%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_debug_report%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>courtney-g</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-12-14

**IP Status**  
No known IP claims.

**Contributors**  
- Courtney Goeltzenleuchter, LunarG

- Dan Ginsburg, Valve

- Jon Ashburn, LunarG

- Mark Lobodzinski, LunarG

## <a href="#_description" class="anchor"></a>Description

Due to the nature of the Vulkan interface, there is very little error
information available to the developer and application. By enabling
optional validation layers and using the `VK_EXT_debug_report`
extension, developers **can** obtain much more detailed feedback on the
application’s use of Vulkan. This extension defines a way for layers and
the implementation to call back to the application for events of
interest to the application.

## <a href="#_new_object_types" class="anchor"></a>New Object Types

- [VkDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportCallbackEXT.html)

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCreateDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDebugReportCallbackEXT.html)

- [vkDebugReportMessageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDebugReportMessageEXT.html)

- [vkDestroyDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyDebugReportCallbackEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateInfo.html):

  - [VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportCallbackCreateInfoEXT.html)

## <a href="#_new_function_pointers" class="anchor"></a>New Function Pointers

- [PFN_vkDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkDebugReportCallbackEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkDebugReportFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportFlagBitsEXT.html)

- [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportObjectTypeEXT.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkDebugReportFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportFlagsEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_DEBUG_REPORT_EXTENSION_NAME`

- `VK_EXT_DEBUG_REPORT_SPEC_VERSION`

- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html):

  - `VK_OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT`

- Extending [VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html):

  - `VK_ERROR_VALIDATION_FAILED_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_DEBUG_REPORT_CALLBACK_CREATE_INFO_EXT`

  - `VK_STRUCTURE_TYPE_DEBUG_REPORT_CREATE_INFO_EXT`

If [Version 1.1](#versions-1.1) is supported:

- Extending
  [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportObjectTypeEXT.html):

  - `VK_DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_EXT`

  - `VK_DEBUG_REPORT_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_EXT`

## <a href="#_examples" class="anchor"></a>Examples

`VK_EXT_debug_report` allows an application to register multiple
callbacks with the validation layers. Some callbacks may log the
information to a file, others may cause a debug break point or other
application defined behavior. An application **can** register callbacks
even when no validation layers are enabled, but they will only be called
for loader and, if implemented, driver events.

To capture events that occur while creating or destroying an instance an
application **can** link a
[VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportCallbackCreateInfoEXT.html)
structure to the `pNext` element of the
[VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateInfo.html) structure given to
[vkCreateInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateInstance.html).

Example uses: Create three callback objects. One will log errors and
warnings to the debug console using Windows `OutputDebugString`. The
second will cause the debugger to break at that callback when an error
happens and the third will log warnings to stdout.

``` c
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

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>In the initial release of the <code>VK_EXT_debug_report</code>
extension, the token
<code>VK_STRUCTURE_TYPE_DEBUG_REPORT_CREATE_INFO_EXT</code> was used.
Starting in version 2 of the extension branch,
<code>VK_STRUCTURE_TYPE_DEBUG_REPORT_CALLBACK_CREATE_INFO_EXT</code> is
used instead for consistency with Vulkan naming rules. The older enum is
still available for backwards compatibility.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>In the initial release of the <code>VK_EXT_debug_report</code>
extension, the token
<code>VK_DEBUG_REPORT_OBJECT_TYPE_DEBUG_REPORT_EXT</code> was used.
Starting in version 8 of the extension branch,
<code>VK_DEBUG_REPORT_OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT_EXT</code>
is used instead for consistency with Vulkan naming rules. The older enum
is still available for backwards compatibility.</p></td>
</tr>
</tbody>
</table>

## <a href="#_issues" class="anchor"></a>Issues

1\) What is the hierarchy / seriousness of the message flags? E.g.
`ERROR` \> `WARN` \> `PERF_WARN` …​

**RESOLVED**: There is no specific hierarchy. Each bit is independent
and should be checked via bitwise AND. For example:

``` c
    if (localFlags & VK_DEBUG_REPORT_ERROR_BIT_EXT) {
        process error message
    }
    if (localFlags & VK_DEBUG_REPORT_DEBUG_BIT_EXT) {
        process debug message
    }
```

The validation layers do use them in a hierarchical way (`ERROR` \>
`WARN` \> `PERF`, `WARN` \> `DEBUG` \> `INFO`) and they (at least at the
time of this writing) only set one bit at a time. But it is not a
requirement of this extension.

It is possible that a layer may intercept and change, or augment the
flags with extension values the application’s debug report handler may
not be familiar with, so it is important to treat each flag
independently.

2\) Should there be a VU requiring
[VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportCallbackCreateInfoEXT.html)::`flags`
to be non-zero?

**RESOLVED**: It may not be very useful, but we do not need VU statement
requiring the `VkDebugReportCallbackCreateInfoEXT`::`msgFlags` at
create-time to be non-zero. One can imagine that apps may prefer it as
it allows them to set the mask as desired - including nothing - at
runtime without having to check.

3\) What is the difference between `VK_DEBUG_REPORT_DEBUG_BIT_EXT` and
`VK_DEBUG_REPORT_INFORMATION_BIT_EXT`?

**RESOLVED**: `VK_DEBUG_REPORT_DEBUG_BIT_EXT` specifies information that
could be useful debugging the Vulkan implementation itself.

4\) How do you compare handles returned by the debug_report callback to
the application’s handles?

**RESOLVED**: Due to the different nature of dispatchable and
nondispatchable handles there is no generic way (that we know of) that
works for common compilers with 32bit, 64bit, C and C++. We recommend
applications use the same cast that the validation layers use:

\+

``` c
reinterpret_cast<uint64_t &>(dispatchableHandle)
(uint64_t)(nondispatchableHandle)
```

\+ This does require that the app treat dispatchable and nondispatchable
handles differently.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2015-05-20 (Courtney Goetzenleuchter)

  - Initial draft, based on LunarG KHR spec, other KHR specs

- Revision 2, 2016-02-16 (Courtney Goetzenleuchter)

  - Update usage, documentation

- Revision 3, 2016-06-14 (Courtney Goetzenleuchter)

  - Update VK_EXT_DEBUG_REPORT_SPEC_VERSION to indicate added support
    for vkCreateInstance and vkDestroyInstance

- Revision 4, 2016-12-08 (Mark Lobodzinski)

  - Added Display_KHR, DisplayModeKHR extension objects

  - Added ObjectTable_NVX, IndirectCommandsLayout_NVX extension objects

  - Bumped spec revision

  - Retroactively added version history

- Revision 5, 2017-01-31 (Baldur Karlsson)

  - Moved definition of
    [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportObjectTypeEXT.html) from
    debug marker chapter

- Revision 6, 2017-01-31 (Baldur Karlsson)

  - Added VK_DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_KHR_EXT

- Revision 7, 2017-04-20 (Courtney Goeltzenleuchter)

  - Clarify wording and address questions from developers.

- Revision 8, 2017-04-21 (Courtney Goeltzenleuchter)

  - Remove unused enum VkDebugReportErrorEXT

- Revision 9, 2017-09-12 (Tobias Hector)

  - Added interactions with Vulkan 1.1

- Revision 10, 2020-12-14 (Courtney Goetzenleuchter)

  - Add issue 4 discussing matching handles returned by the extension,
    based on suggestion in public issue 368.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_debug_report"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
