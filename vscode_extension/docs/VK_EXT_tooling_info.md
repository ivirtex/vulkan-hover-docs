# VK\_EXT\_tooling\_info(3) Manual Page

## Name

VK\_EXT\_tooling\_info - device extension



## [](#_registered_extension_number)Registered Extension Number

246

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_api_interactions)API Interactions

- Interacts with VK\_EXT\_debug\_marker
- Interacts with VK\_EXT\_debug\_report
- Interacts with VK\_EXT\_debug\_utils

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.3](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.3-promotions)

## [](#_contact)Contact

- Tobias Hector [\[GitHub\]tobski](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_tooling_info%5D%20%40tobski%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_tooling_info%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-11-05

**Contributors**

- Rolando Caloca
- Matthaeus Chajdas
- Baldur Karlsson
- Daniel Rakos

## [](#_description)Description

When an error occurs during application development, a common question is "What tools are actually running right now?" This extension adds the ability to query that information directly from the Vulkan implementation.

Outdated versions of one tool might not play nicely with another, or perhaps a tool is not actually running when it should have been. Trying to figure that out can cause headaches as it is necessary to consult each known tool to figure out what is going on — in some cases the tool might not even be known.

Typically, the expectation is that developers will simply print out this information for visual inspection when an issue occurs, however a small amount of semantic information about what the tool is doing is provided to help identify it programmatically. For example, if the advertised limits or features of an implementation are unexpected, is there a tool active which modifies these limits? Or if an application is providing debug markers, but the implementation is not actually doing anything with that information, this can quickly point that out.

## [](#_new_commands)New Commands

- [vkGetPhysicalDeviceToolPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceToolPropertiesEXT.html)

## [](#_new_structures)New Structures

- [VkPhysicalDeviceToolPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceToolPropertiesEXT.html)

## [](#_new_enums)New Enums

- [VkToolPurposeFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkToolPurposeFlagBitsEXT.html)

## [](#_new_bitmasks)New Bitmasks

- [VkToolPurposeFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkToolPurposeFlagsEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_TOOLING_INFO_EXTENSION_NAME`
- `VK_EXT_TOOLING_INFO_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TOOL_PROPERTIES_EXT`
- Extending [VkToolPurposeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkToolPurposeFlagBits.html):
  
  - `VK_TOOL_PURPOSE_ADDITIONAL_FEATURES_BIT_EXT`
  - `VK_TOOL_PURPOSE_MODIFYING_FEATURES_BIT_EXT`
  - `VK_TOOL_PURPOSE_PROFILING_BIT_EXT`
  - `VK_TOOL_PURPOSE_TRACING_BIT_EXT`
  - `VK_TOOL_PURPOSE_VALIDATION_BIT_EXT`

If [VK\_EXT\_debug\_marker](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_marker.html) is supported:

- Extending [VkToolPurposeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkToolPurposeFlagBits.html):
  
  - `VK_TOOL_PURPOSE_DEBUG_MARKERS_BIT_EXT`

If [VK\_EXT\_debug\_report](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_report.html) is supported:

- Extending [VkToolPurposeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkToolPurposeFlagBits.html):
  
  - `VK_TOOL_PURPOSE_DEBUG_REPORTING_BIT_EXT`

If [VK\_EXT\_debug\_utils](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_utils.html) is supported:

- Extending [VkToolPurposeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkToolPurposeFlagBits.html):
  
  - `VK_TOOL_PURPOSE_DEBUG_MARKERS_BIT_EXT`
  - `VK_TOOL_PURPOSE_DEBUG_REPORTING_BIT_EXT`

## [](#_promotion_to_vulkan_1_3)Promotion to Vulkan 1.3

Vulkan APIs in this extension are included in core Vulkan 1.3, with the EXT suffix omitted. External interactions defined by this extension, such as SPIR-V token names, retain their original names. The original Vulkan API names are still available as aliases of the core functionality.

## [](#_examples)Examples

Printing Tool Information

```
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

## [](#_issues)Issues

1\) Why is this information separate from the layer mechanism?

Some tooling may be built into a driver, or be part of the Vulkan loader etc. Tying this information directly to layers would have been awkward at best.

## [](#_version_history)Version History

- Revision 1, 2018-11-05 (Tobias Hector)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_tooling_info).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0