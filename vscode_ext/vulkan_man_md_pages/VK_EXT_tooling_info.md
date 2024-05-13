# VK_EXT_tooling_info(3) Manual Page

## Name

VK_EXT_tooling_info - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

246

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_EXT_debug_marker

- Interacts with VK_EXT_debug_report

- Interacts with VK_EXT_debug_utils

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.3-promotions"
  target="_blank" rel="noopener">Vulkan 1.3</a>

## <a href="#_contact" class="anchor"></a>Contact

- Tobias Hector <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_tooling_info%5D%20@tobski%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_tooling_info%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>tobski</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-11-05

**Contributors**  
- Rolando Caloca

- Matthaeus Chajdas

- Baldur Karlsson

- Daniel Rakos

## <a href="#_description" class="anchor"></a>Description

When an error occurs during application development, a common question
is "What tools are actually running right now?" This extension adds the
ability to query that information directly from the Vulkan
implementation.

Outdated versions of one tool might not play nicely with another, or
perhaps a tool is not actually running when it should have been. Trying
to figure that out can cause headaches as it is necessary to consult
each known tool to figure out what is going on — in some cases the tool
might not even be known.

Typically, the expectation is that developers will simply print out this
information for visual inspection when an issue occurs, however a small
amount of semantic information about what the tool is doing is provided
to help identify it programmatically. For example, if the advertised
limits or features of an implementation are unexpected, is there a tool
active which modifies these limits? Or if an application is providing
debug markers, but the implementation is not actually doing anything
with that information, this can quickly point that out.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetPhysicalDeviceToolPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceToolPropertiesEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkPhysicalDeviceToolPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceToolPropertiesEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkToolPurposeFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkToolPurposeFlagBitsEXT.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkToolPurposeFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkToolPurposeFlagsEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_TOOLING_INFO_EXTENSION_NAME`

- `VK_EXT_TOOLING_INFO_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TOOL_PROPERTIES_EXT`

If [VK_EXT_debug_marker](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_marker.html) is supported:

- Extending [VkToolPurposeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkToolPurposeFlagBits.html):

  - `VK_TOOL_PURPOSE_DEBUG_MARKERS_BIT_EXT`

If [VK_EXT_debug_report](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_report.html) is supported:

- Extending [VkToolPurposeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkToolPurposeFlagBits.html):

  - `VK_TOOL_PURPOSE_DEBUG_REPORTING_BIT_EXT`

If [VK_EXT_debug_utils](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_utils.html) is supported:

- Extending [VkToolPurposeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkToolPurposeFlagBits.html):

  - `VK_TOOL_PURPOSE_DEBUG_MARKERS_BIT_EXT`

  - `VK_TOOL_PURPOSE_DEBUG_REPORTING_BIT_EXT`

## <a href="#_promotion_to_vulkan_1_3" class="anchor"></a>Promotion to Vulkan 1.3

Functionality in this extension is included in core Vulkan 1.3, with the
EXT suffix omitted. The original type, enum and command names are still
available as aliases of the core functionality.

## <a href="#_examples" class="anchor"></a>Examples

Printing Tool Information

``` c
uint32_t num_tools;
VkPhysicalDeviceToolPropertiesEXT *pToolProperties;
vkGetPhysicalDeviceToolPropertiesEXT(physicalDevice, &num_tools, NULL);

pToolProperties = (VkPhysicalDeviceToolPropertiesEXT*)malloc(sizeof(VkPhysicalDeviceToolPropertiesEXT) * num_tools);

vkGetPhysicalDeviceToolPropertiesEXT(physicalDevice, &num_tools, pToolProperties);

for (int i = 0; i < num_tools; ++i) {
    printf("%s:\n", pToolProperties[i].name);
    printf("Version:\n");
    printf("%s:\n", pToolProperties[i].version);
    printf("Description:\n");
    printf("\t%s\n", pToolProperties[i].description);
    printf("Purposes:\n");
    printf("\t%s\n", VkToolPurposeFlagBitsEXT_to_string(pToolProperties[i].purposes));
    if (strnlen_s(pToolProperties[i].layer,VK_MAX_EXTENSION_NAME_SIZE) > 0) {
        printf("Corresponding Layer:\n");
        printf("\t%s\n", pToolProperties[i].layer);
    }
}
```

## <a href="#_issues" class="anchor"></a>Issues

1\) Why is this information separate from the layer mechanism?

Some tooling may be built into a driver, or be part of the Vulkan loader
etc. Tying this information directly to layers would have been awkward
at best.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-11-05 (Tobias Hector)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_tooling_info"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
