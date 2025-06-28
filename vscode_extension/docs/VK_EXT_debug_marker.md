# VK\_EXT\_debug\_marker(3) Manual Page

## Name

VK\_EXT\_debug\_marker - device extension



## [](#_registered_extension_number)Registered Extension Number

23

## [](#_revision)Revision

4

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_EXT\_debug\_report](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_report.html)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [VK\_EXT\_debug\_utils](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_utils.html) extension

## [](#_special_use)Special Use

- [Debugging tools](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Baldur Karlsson [\[GitHub\]baldurk](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_debug_marker%5D%20%40baldurk%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_debug_marker%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-01-31

**IP Status**

No known IP claims.

**Contributors**

- Baldur Karlsson
- Dan Ginsburg, Valve
- Jon Ashburn, LunarG
- Kyle Spagnoli, NVIDIA

## [](#_description)Description

The `VK_EXT_debug_marker` extension is a device extension. It introduces concepts of object naming and tagging, for better tracking of Vulkan objects, as well as additional commands for recording annotations of named sections of a workload to aid organization and offline analysis in external tools.

## [](#_new_commands)New Commands

- [vkCmdDebugMarkerBeginEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDebugMarkerBeginEXT.html)
- [vkCmdDebugMarkerEndEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDebugMarkerEndEXT.html)
- [vkCmdDebugMarkerInsertEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDebugMarkerInsertEXT.html)
- [vkDebugMarkerSetObjectNameEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDebugMarkerSetObjectNameEXT.html)
- [vkDebugMarkerSetObjectTagEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDebugMarkerSetObjectTagEXT.html)

## [](#_new_structures)New Structures

- [VkDebugMarkerMarkerInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugMarkerMarkerInfoEXT.html)
- [VkDebugMarkerObjectNameInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugMarkerObjectNameInfoEXT.html)
- [VkDebugMarkerObjectTagInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugMarkerObjectTagInfoEXT.html)

## [](#_new_enums)New Enums

- [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportObjectTypeEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_DEBUG_MARKER_EXTENSION_NAME`
- `VK_EXT_DEBUG_MARKER_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DEBUG_MARKER_MARKER_INFO_EXT`
  - `VK_STRUCTURE_TYPE_DEBUG_MARKER_OBJECT_NAME_INFO_EXT`
  - `VK_STRUCTURE_TYPE_DEBUG_MARKER_OBJECT_TAG_INFO_EXT`

## [](#_examples)Examples

**Example 1**

Associate a name with an image, for easier debugging in external tools or with validation layers that can print a friendly name when referring to objects in error messages.

```c++
    extern VkDevice device;
    extern VkImage image;

    // Must call extension functions through a function pointer:
    PFN_vkDebugMarkerSetObjectNameEXT pfnDebugMarkerSetObjectNameEXT = (PFN_vkDebugMarkerSetObjectNameEXT)vkGetDeviceProcAddr(device, "vkDebugMarkerSetObjectNameEXT");

    // Set a name on the image
    const VkDebugMarkerObjectNameInfoEXT imageNameInfo =
    {
        .sType = VK_STRUCTURE_TYPE_DEBUG_MARKER_OBJECT_NAME_INFO_EXT,
        .pNext = NULL,
        .objectType = VK_DEBUG_REPORT_OBJECT_TYPE_IMAGE_EXT,
        .object = (uint64_t)image,
        .pObjectName = "Brick Diffuse Texture",
    };

    pfnDebugMarkerSetObjectNameEXT(device, &imageNameInfo);

    // A subsequent error might print:
    //   Image 'Brick Diffuse Texture' (0xc0dec0dedeadbeef) is used in a
    //   command buffer with no memory bound to it.
```

**Example 2**

Annotating regions of a workload with naming information so that offline analysis tools can display a more usable visualization of the commands submitted.

```c++
    extern VkDevice device;
    extern VkCommandBuffer commandBuffer;

    // Must call extension functions through a function pointer:
    PFN_vkCmdDebugMarkerBeginEXT pfnCmdDebugMarkerBeginEXT = (PFN_vkCmdDebugMarkerBeginEXT)vkGetDeviceProcAddr(device, "vkCmdDebugMarkerBeginEXT");
    PFN_vkCmdDebugMarkerEndEXT pfnCmdDebugMarkerEndEXT = (PFN_vkCmdDebugMarkerEndEXT)vkGetDeviceProcAddr(device, "vkCmdDebugMarkerEndEXT");
    PFN_vkCmdDebugMarkerInsertEXT pfnCmdDebugMarkerInsertEXT = (PFN_vkCmdDebugMarkerInsertEXT)vkGetDeviceProcAddr(device, "vkCmdDebugMarkerInsertEXT");

    // Describe the area being rendered
    const VkDebugMarkerMarkerInfoEXT houseMarker =
    {
        .sType = VK_STRUCTURE_TYPE_DEBUG_MARKER_MARKER_INFO_EXT,
        .pNext = NULL,
        .pMarkerName = "Brick House",
        .color = { 1.0f, 0.0f, 0.0f, 1.0f },
    };

    // Start an annotated group of calls under the 'Brick House' name
    pfnCmdDebugMarkerBeginEXT(commandBuffer, &houseMarker);
    {
        // A mutable structure for each part being rendered
        VkDebugMarkerMarkerInfoEXT housePartMarker =
        {
            .sType = VK_STRUCTURE_TYPE_DEBUG_MARKER_MARKER_INFO_EXT,
            .pNext = NULL,
            .pMarkerName = NULL,
            .color = { 0.0f, 0.0f, 0.0f, 0.0f },
        };

        // Set the name and insert the marker
        housePartMarker.pMarkerName = "Walls";
        pfnCmdDebugMarkerInsertEXT(commandBuffer, &housePartMarker);

        // Insert the drawcall for the walls
        vkCmdDrawIndexed(commandBuffer, 1000, 1, 0, 0, 0);

        // Insert a recursive region for two sets of windows
        housePartMarker.pMarkerName = "Windows";
        pfnCmdDebugMarkerBeginEXT(commandBuffer, &housePartMarker);
        {
            vkCmdDrawIndexed(commandBuffer, 75, 6, 1000, 0, 0);
            vkCmdDrawIndexed(commandBuffer, 100, 2, 1450, 0, 0);
        }
        pfnCmdDebugMarkerEndEXT(commandBuffer);

        housePartMarker.pMarkerName = "Front Door";
        pfnCmdDebugMarkerInsertEXT(commandBuffer, &housePartMarker);

        vkCmdDrawIndexed(commandBuffer, 350, 1, 1650, 0, 0);

        housePartMarker.pMarkerName = "Roof";
        pfnCmdDebugMarkerInsertEXT(commandBuffer, &housePartMarker);

        vkCmdDrawIndexed(commandBuffer, 500, 1, 2000, 0, 0);
    }
    // End the house annotation started above
    pfnCmdDebugMarkerEndEXT(commandBuffer);
```

## [](#_issues)Issues

1\) Should the tag or name for an object be specified using the `pNext` parameter in the object’s `Vk*CreateInfo` structure?

**RESOLVED**: No. While this fits with other Vulkan patterns and would allow more type safety and future proofing against future objects, it has notable downsides. In particular passing the name at `Vk*CreateInfo` time does not allow renaming, prevents late binding of naming information, and does not allow naming of implicitly created objects such as queues and swapchain images.

2\) Should the command annotation functions [vkCmdDebugMarkerBeginEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDebugMarkerBeginEXT.html) and [vkCmdDebugMarkerEndEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDebugMarkerEndEXT.html) support the ability to specify a color?

**RESOLVED**: Yes. The functions have been expanded to take an optional color which can be used at will by implementations consuming the command buffer annotations in their visualization.

3\) Should the functions added in this extension accept an extensible structure as their parameter for a more flexible API, as opposed to direct function parameters? If so, which functions?

**RESOLVED**: Yes. All functions have been modified to take a structure type with extensible `pNext` pointer, to allow future extensions to add additional annotation information in the same commands.

## [](#_version_history)Version History

- Revision 1, 2016-02-24 (Baldur Karlsson)
  
  - Initial draft, based on LunarG marker spec
- Revision 2, 2016-02-26 (Baldur Karlsson)
  
  - Renamed Dbg to DebugMarker in function names
  - Allow markers in secondary command buffers under certain circumstances
  - Minor language tweaks and edits
- Revision 3, 2016-04-23 (Baldur Karlsson)
  
  - Reorganize spec layout to closer match desired organization
  - Added optional color to markers (both regions and inserted labels)
  - Changed functions to take extensible structs instead of direct function parameters
- Revision 4, 2017-01-31 (Baldur Karlsson)
  
  - Added explicit dependency on VK\_EXT\_debug\_report
  - Moved definition of [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportObjectTypeEXT.html) to debug report chapter.
  - Fixed typo in dates in revision history

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_debug_marker)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0